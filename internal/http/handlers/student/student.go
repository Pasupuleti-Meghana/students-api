package student

import (
	"Pasupuleti-Meghana/students-api/internal/types"
	"Pasupuleti-Meghana/students-api/internal/utils"
	"encoding/json"
	"strconv"

	// "log"

	// "errors"
	// "io"
	"log/slog"
	"net/http"

	// "fmt"
	"Pasupuleti-Meghana/students-api/internal/storage"
	"fmt"

	"github.com/go-playground/validator/v10"
)
   
func New(storage storage.Storage) http.HandlerFunc {
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
		if err != nil {
			validationErrs := err.(validator.ValidationErrors)
			response.WriteJsonResponse(w, http.StatusBadRequest,response.ValidationErrorResponse(validationErrs))
			return
		}
 

		lastId, err := storage.CreateStudent(
			student.Name,
			student.Email,
			student.Age,
		)

		if err != nil {
			response.WriteJsonResponse(w, http.StatusInternalServerError, err.Error())
		}

		slog.Info("Student created succesfully..", slog.String("userId", fmt.Sprint(lastId)))

		response.WriteJsonResponse(w, http.StatusOK, map[string]int64{ "id": lastId})

	}
}

func GetById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		slog.Info("getting a student",slog.String("id",id))

		intId, err:= strconv.ParseInt(id, 10, 64)
		if err != nil {
			slog.Error("Erro in first err",err)
			response.WriteJsonResponse(w, http.StatusBadRequest, response.GeneralErrorResponse(err))
			return
		}

		student, err := storage.GetStudentById(intId)
		if err != nil {
			slog.Error("Error getting user",slog.String("id",id))
			response.WriteJsonResponse(w, http.StatusInternalServerError, response.GeneralErrorResponse(err))
			return
		}

		response.WriteJsonResponse(w, http.StatusOK, student)
	}
}

func GetList(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Getting Users")

		students, err := storage.GetStudentList()
		if err != nil {
			response.WriteJsonResponse(w, http.StatusInternalServerError, err)
		}

		response.WriteJsonResponse(w, http.StatusOK, students)
	}
}