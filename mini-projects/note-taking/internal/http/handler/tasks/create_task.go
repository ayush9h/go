package create_task

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/ayush9h/internal/types"
	"github.com/go-playground/validator/v10"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/tasks/create", createTaskHandler)
}

func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	var request types.CreateTaskRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, "Invalid request json", http.StatusBadRequest)
		return
	}
	validate := validator.New()

	err = validate.Struct(request)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		http.Error(w, fmt.Sprintf("Validation error %s", errors), http.StatusBadRequest)
		return
	}
	response := types.CreateTaskResponse{Message: "Note creatd", Data: request}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response_err := json.NewEncoder(w).Encode(response)
	if response_err != nil {
		slog.Error("Failed to encode response", "error", response_err.Error())
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	slog.Info("Task Created Successfully")
}
