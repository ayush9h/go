package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to Note Taking Page.")
	})

	slog.Info("Server started on port 8000")
	err := http.ListenAndServe(":8000", mux)

	if err != nil {
		slog.Error("failure to start the server due to", slog.String("error", err.Error()))
	}
}
