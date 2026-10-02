package core_transport_http_utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

func GetJson[T any](r *http.Request) (T, error) {
	var req T
	var reqNil T

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return reqNil, fmt.Errorf("Failed to parse: %w", err)
	}

	return req, nil
}

func GetAndFilterJson[T any](validator *validator.Validate, r *http.Request) (T, error) {
	var req T
	var reqNil T

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return reqNil, fmt.Errorf("Failed to parse: %w", err)
	}

	if err := validator.Struct(req); err != nil {
		return reqNil, fmt.Errorf("The structure failed validation: %w", err)
	}

	return req, nil
}

func GetToken(r *http.Request, key string) (*string, error) {
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		return nil, fmt.Errorf("authorization header is required")
	}

	const prefix = "Bearer "

	if !strings.HasPrefix(authHeader, prefix) {
		return nil, fmt.Errorf("invalid authorization header")
	}

	token := strings.TrimPrefix(authHeader, prefix)

	if token == "" {
		return nil, fmt.Errorf("token is required")
	}

	return &token, nil
}
