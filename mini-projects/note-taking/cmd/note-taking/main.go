package main

import (
	"fmt"
	"log/slog"
	"net/http"

	tasks "github.com/ayush9h/internal/http/handler/tasks"
)

func registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Welcome to the note taking page")
	})

	// Register create task routes
	tasks.RegisterRoutes(mux)
}

func main() {

	mux := http.NewServeMux()

	registerRoutes(mux)

	slog.Info("Server started on port 8000")
	err := http.ListenAndServe(":8000", mux)

	if err != nil {
		slog.Error("failure to start the server due to", slog.String("error", err.Error()))
	}
}
