package main

import (
	"fmt"
	"net/http"
	"os"
	"project_factions/internal/logging"
	"project_factions/internal/web"

	"github.com/go-chi/chi/v5"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logging.Info("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

func main() {
	// LOGGING
	logging.Init("", true)

	// HANDLERS
	r := chi.NewRouter()
	r.Use(loggingMiddleware)

	h := web.NewHandler()
	r.Get("/", h.HandleIndex)

	r.Post("/generate", h.HandleGenerate)

	// SERVER
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if port == ":" {
		port = ":8080" // Default port to listen on
	}
	logging.Info("Server started at port %s", port)

	http.ListenAndServe(port, r)
}
