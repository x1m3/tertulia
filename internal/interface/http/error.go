package http

import (
	"encoding/json"
	"errors"
	"net/http"

	apiErrors "github.com/x1m3/tertulia/internal/interface/http/errors"
)

// ErrorHandlerFunc is an error adapter for the API. It is used to standardize errors happening in the generated code api handlers
func ErrorHandlerFunc(w http.ResponseWriter, _ *http.Request, err error) {
	var invalidParamFormatError *InvalidParamFormatError
	switch {
	case errors.As(err, &invalidParamFormatError):
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{"message": err.Error()})
	default:
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

// RequestErrorHandlerFunc is a Request Error Handler that can be injected in oapi-codegen to handler errors in requests
func RequestErrorHandlerFunc(w http.ResponseWriter, _ *http.Request, err error) {
	http.Error(w, err.Error(), http.StatusBadRequest)
}

// ResponseErrorHandlerFunc is a Response Error Handler that can be injected in oapi-codegen to handler errors in requests
// We use it to create custom responses to some errors that may occur, like an authentication error.
func ResponseErrorHandlerFunc(w http.ResponseWriter, _ *http.Request, err error) {
	w.Header().Add("Content-Type", "application/json")
	var authError apiErrors.AuthError
	switch {
	case errors.As(err, &authError):
		w.WriteHeader(http.StatusUnauthorized)
		w.Header().Add("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		_, _ = w.Write([]byte("\"Unauthorized\""))
	default:
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
	}
}
