package responses

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
)

type Response struct {
	Status string
	Error  string
}

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

func WriteJson(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func GeneralError(err error) Response {
	return Response{
		Status: StatusError,
		Error:  err.Error(),
	}
}

func ValidationError(err validator.validationErrors) Response{
	var errMsgs []string

	for _, err := range errs{
			switch err.ActualTag(){

					case "required":
						errMsgs = append(errMsgs, fmt.Sprintf("field %s is required field",err.Field()))

					default:
						errMsgs = append(errMsgs, fmt.Sprintf("field %s is invalid field",err.Field()))
			}
		}
	}
	return Response{
		Status: StatusError,
		Error: strings.Join(errMsgs,", ")
	}
}
