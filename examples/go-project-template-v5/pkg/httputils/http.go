package httputils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

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

func PathValueI64(r *http.Request, name string) (int64, error) {
	pathValue := r.PathValue(name)
	if pathValue == "" {
		return 0, fmt.Errorf("empty path value for name: %s", name)
	}
	return strconv.ParseInt(pathValue, 10, 64)
}
