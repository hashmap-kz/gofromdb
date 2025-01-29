package middlewares

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"go-project-template-v5/config"

	"github.com/golang-jwt/jwt"
)

const (
	HTTPHeaderAuthorization = "Authorization"
	TokenTypeHintAccess     = "access_token"
	TokenTypeHintRPT        = "requesting_party_token"
)

func AuthorizeMiddleware(nextHandler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := config.Cfg().Auth

		c := &Client{
			URL:          auth.URL,
			Realm:        auth.Realm,
			ClientID:     auth.ClientId,
			ClientSecret: auth.ClientSecret,
		}

		_, err := c.Introspect(r.Header.Get(HTTPHeaderAuthorization), TokenTypeHintRPT)
		if err != nil {
			w.WriteHeader(401)
			return
		}

		nextHandler.ServeHTTP(w, r)
	}
}

type Introspect struct {
	Active       bool   `json:"active"`
	Username     string `json:"username,omitempty"`
	Email        string `json:"email,omitempty"`
	ClientID     string `json:"client_id,omitempty"`
	Sub          string `json:"sub,omitempty"`
	Exp          int64  `json:"exp,omitempty"`
	Iat          int64  `json:"iat,omitempty"`
	Jti          string `json:"jti,omitempty"`
	Iss          string `json:"iss,omitempty"`
	Typ          string `json:"typ,omitempty"`
	SessionState string `json:"session_state,omitempty"`
	Name         string `json:"name,omitempty"`
	Scope        string `json:"scope,omitempty"`
}

type Client struct {
	URL          string
	Realm        string
	ClientID     string
	ClientSecret string
}

func (c *Client) formatToken(token string) string {
	if strings.HasPrefix(token, "Bearer ") {
		return token[7:]
	}
	return token
}

// https://www.oauth.com/oauth2-servers/token-introspection-endpoint/
func (c *Client) Introspect(token, tokenTypeHint string) (*Introspect, error) {
	formattedToken := c.formatToken(token)
	if formattedToken == "" {
		return nil, errors.New("no token provided")
	}

	formData := url.Values{
		"client_id":       {c.ClientID},
		"client_secret":   {c.ClientSecret},
		"token":           {formattedToken},
		"token_type_hint": {tokenTypeHint},
	}
	req, err := http.NewRequest("POST", c.URL+"/realms/"+c.Realm+"/protocol/openid-connect/token/introspect", strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		return nil, fmt.Errorf("error requesting token introspection: %w", err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		if resp == nil {
			return nil, errors.New("failed to introspect token")
		}
		body, err := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to introspect token. Reason: %s, %w", string(body), err)
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.New("failed to introspect token")
		}
		return nil, fmt.Errorf("failed to introspect token. Reason: %s", string(body))
	}

	var introspect Introspect
	err = json.NewDecoder(resp.Body).Decode(&introspect)
	if err != nil {
		return nil, fmt.Errorf("error decoding introspection response: %w", err)
	}
	if !introspect.Active {
		return nil, fmt.Errorf("token is not active")
	}

	return &introspect, nil
}

// parse JWT token with custom attributes/claims

type claims struct {
	jwt.StandardClaims

	Audience       any                            `json:"aud,omitempty"`
	Subject        string                         `json:"sub"`
	ResourceAccess map[string]map[string][]string `json:"resource_access"`

	// claims from custom protocol mapper

	SpecRoles []struct {
		Name        string              `json:"name,omitempty"`
		Description string              `json:"description,omitempty"`
		Attributes  map[string][]string `json:"attributes,omitempty"`
	} `json:"specRoles,omitempty"`

	SpecGroups []struct {
		Name       string              `json:"name,omitempty"`
		ID         string              `json:"id,omitempty"`
		ParentID   string              `json:"parentId,omitempty"`
		Attributes map[string][]string `json:"attributes,omitempty"`
	} `json:"specGroups,omitempty"`

	SpecAttributes map[string][]string `json:"specAttributes,omitempty"`

	SpecAuthorities []string `json:"specAuthorities,omitempty"`
}

func parseTokenUnverified(tokenStr string) (*jwt.Token, *claims, error) {
	tokenClaims := &claims{}
	jwtToken, _, err := new(jwt.Parser).ParseUnverified(tokenStr, tokenClaims)
	if err != nil {
		return nil, nil, err
	}
	return jwtToken, tokenClaims, nil
}
