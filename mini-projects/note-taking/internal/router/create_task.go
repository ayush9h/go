package create_task

import (
	"fmt"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/v1/tasks/create", createTaskHandler)
}

func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Task Creation Page")
}
