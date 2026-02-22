package response

import (
	"encoding/json"
	"net/http"
	"strings"
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status      string `json:"status"`
	StatusError string `json:"status_error"`
}

const (
	StatusOk    = "OK"
	StatusError = "Error"
)

func WriteJsonResponse(w http.ResponseWriter, statusCode int, data interface{}) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(data)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"Failed to encode response"}`))
	}
}

func GeneralErrorResponse(err error) Response {
	return Response{
		Status:      StatusError,
		StatusError: err.Error(),
	}
}

func ValidationErrorResponse(errs validator.ValidationErrors) Response {
	var errorMsg []string 

	for _, err := range errs {
		if err.Tag() == "required" {
			errorMsg = append(errorMsg, err.Field() + " is required")
		}
	}
	return Response {
		Status : StatusError,
		StatusError: "Validation failed: " +  strings.Join(errorMsg, ", "),
	}
}
