package httputils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
)

var uuidRegex = regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$`)

type ErrorDetail struct {
	Field string `json:"field"`
	Issue string `json:"issue"`
}

type ErrorResponse struct {
	Message string        `json:"message,omitempty"`
	Details []ErrorDetail `json:"details,omitempty"`
}

func ReadJSON(r *http.Request, v any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}
	return json.NewDecoder(r.Body).Decode(v)
}

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func PathValueUUID(r *http.Request, name string) (string, error) {
	pathValue := r.PathValue(name)
	if pathValue == "" {
		return "", fmt.Errorf("empty path value for name: %s", name)
	}
	if uuidRegex.MatchString(pathValue) {
		return pathValue, nil
	}
	return "", fmt.Errorf("path value with name=%s and value=%s is not an UUID", name, pathValue)
}

func PathValueI32(r *http.Request, name string) (int, error) {
	pathValue := r.PathValue(name)
	if pathValue == "" {
		return 0, fmt.Errorf("empty path value for name: %s", name)
	}
	return strconv.Atoi(pathValue)
}

func PathValueI64(r *http.Request, name string) (int64, error) {
	pathValue := r.PathValue(name)
	if pathValue == "" {
		return 0, fmt.Errorf("empty path value for name: %s", name)
	}
	return strconv.ParseInt(pathValue, 10, 64)
}
