package students

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/ayush9h/students-api/internal/types"
	"github.com/ayush9h/students-api/internal/utils/responses"
	"github.com/go-playground/validator/v10"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			responses.WriteJson(w, http.StatusBadRequest, responses.GeneralError(err))
			return
		}

		//request validation

		if err := validator.New().Struct(student); err != nil {
			validateErrors := err.(validator.ValidationErrors)
			responses.WriteJson(w, http.StatusBadRequest, responses.ValidationError(validateErrors))
			return
		}

		responses.WriteJson(w, http.StatusCreated, map[string]string{"success": "OK"})
	}
}
