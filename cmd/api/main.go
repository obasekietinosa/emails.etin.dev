package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/tinrab/emails/internal/database"
	"github.com/tinrab/emails/internal/handlers"
	"github.com/tinrab/emails/internal/middleware"
)

func main() {
	// Setup DB
	if err := database.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Setup Router
	r := chi.NewRouter()
	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Recoverer)

	authHandler := &handlers.AuthHandler{}

	r.Post("/register", authHandler.Register)
	r.Post("/login", authHandler.Login)
	r.Post("/trigger/{token}", handlers.TriggerEmail)

	r.Group(func(r chi.Router) {
		r.Use(middleware.JWTAuth)
		r.Post("/api-keys", handlers.GenerateAPIKey)

		r.Get("/templates", handlers.ListTemplates)
		r.Post("/templates", handlers.CreateTemplate)
		r.Get("/templates/{id}", handlers.GetTemplate)
		r.Put("/templates/{id}", handlers.UpdateTemplate)
		r.Delete("/templates/{id}", handlers.DeleteTemplate)
	})

	r.Group(func(r chi.Router) {
		r.Use(middleware.APIKeyAuth)
		r.Post("/send", handlers.SendEmail)
	})

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
