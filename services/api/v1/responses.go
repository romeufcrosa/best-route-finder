package v1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// commonly used errors
var (
	ErrNotFound      = errors.New("route does not exist")
	ErrCartNoContent = errors.New("no cart data provided")
)

// Params params to enter in controller
type Params struct{}

// ResultError contains the error code and message
type ResultError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Result the final result for a given message
type Result struct {
	Error  *ResultError    `json:"error,omitempty"`
	Result json.RawMessage `json:"result,omitempty"`
}

// Jsonable an entity that returns a JSON representation of itself
type Jsonable interface {
	ToJSON() (json.RawMessage, error)
}

// Error sends an error
func Error(ctx context.Context, w http.ResponseWriter, err error) {
	result, _ := json.Marshal(Result{
		Error: &ResultError{
			Code:    1,
			Message: err.Error(),
		},
	})

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(result))
}

// Response sends a response
func Response(ctx context.Context, w http.ResponseWriter, result Jsonable) {
	data, err := result.ToJSON()
	if err != nil {
		Error(ctx, w, err)
		return
	}

	response, _ := json.Marshal(data)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(response))
}
