package student

import (
	"Pasupuleti-Meghana/students-api/internal/types"
	"Pasupuleti-Meghana/students-api/internal/utils"
	"encoding/json"
	// "errors"
	// "io"
	"log/slog"
	"net/http"
	// "fmt"
	"github.com/go-playground/validator/v10"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		slog.Info("creating a student...")

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)

		// if errors.Is(err, io.EOF) {
		// 	response.WriteJsonResponse(w, http.StatusBadRequest, err.Error())
		// 	return 
		// }

		// if err != nil {
		// 	response.WriteJsonResponse(w, http.StatusBadRequest, err.Error())
		// 	return
		// }	

		if err != nil {
			response.WriteJsonResponse(w, http.StatusBadRequest, response.GeneralErrorResponse(err))
			return
		}


		// validate the request 

		err = validator.New().Struct(student)
		validationErrs := err.(validator.ValidationErrors)
		response.WriteJsonResponse(w, http.StatusBadRequest,response.ValidationErrorResponse(validationErrs))
		return 

		response.WriteJsonResponse(w, http.StatusOK, map[string]string{ "message": "Student created successfully"})

	}
}