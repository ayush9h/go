package tasks

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/ayush9h/internal/types"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/tasks/create", createTaskHandler)
	mux.HandleFunc("PUT /api/v1/tasks/{id}", updateTaskHandler)
	mux.HandleFunc("DELETE /api/v1/tasks/{id}", deleteTaskHandler)
}

// Logic for handling creation of tasks
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

	id := uuid.New().String()

	response := types.CreateTaskResponse{Message: "Note creatd", NoteId: id}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response_err := json.NewEncoder(w).Encode(response)
	if response_err != nil {
		slog.Error("Failed to encode response", "error", response_err.Error())
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	// TODO: Store the task in a DB

	slog.Info("Task Created Successfully")

}

// Logic for handling updation of tasks
func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	noteId := r.PathValue("id")

	if noteId == "" {
		http.Error(w, "Note id is null", http.StatusBadRequest)
		return
	}

	var request types.UpdateTaskRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request json", http.StatusBadRequest)
		return
	}
	validate := validator.New()
	err = validate.Struct(request)

	if err != nil {
		errors := err.(validator.ValidationErrors)
		http.Error(w, fmt.Sprintf("error occurred due to %s", errors), http.StatusBadRequest)
		return
	}

	// TODO: checks for notes in the db before updation

	response := types.UpdateTaskResponse{
		Message: "Note updated successfully",
		NoteId:  noteId,
	}

	responseErr := json.NewEncoder(w).Encode(response)
	if responseErr != nil {
		slog.Error("Failed to encode response", "error", responseErr.Error())
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	slog.Info("Task updated")
}

// Logic for handling deletion of tasks
func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	noteId := r.PathValue("id")

	if noteId == "" {
		http.Error(w, "Note id is null", http.StatusBadRequest)
		return
	}

	// TODO: checks for notes in the db before deletion

	response := types.DeleteTaskResponse{
		Message: "Note deleted successfully",
	}

	responseErr := json.NewEncoder(w).Encode(response)
	if responseErr != nil {
		slog.Error("Failed to encode response", "error", responseErr.Error())
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	slog.Info("Task deleted")
}
