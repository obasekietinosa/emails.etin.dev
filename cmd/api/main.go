package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"egogo/internal/config"
	"egogo/internal/database"
	"egogo/internal/handlers"
	"egogo/internal/middleware"
	"egogo/internal/repository"
	"egogo/internal/service"
)

func main() {
	// Setup Logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// Load Config
	cfg := config.Load()

	// Setup DB
	if err := database.InitDB(cfg.DBDSN); err != nil {
		logger.Error("Failed to initialize database", "error", err)
		os.Exit(1)
	}

	// Setup Repository
	repo := repository.NewPostgresRepository(database.DB)

	// Setup Email Service
	var sender service.EmailSender
	if cfg.SMTPHost != "" {
		sender = service.NewSMTPSender(cfg.SMTPHost, cfg.SMTPPort, cfg.SMTPUser, cfg.SMTPPass, cfg.FromEmail)
		logger.Info("Using SMTP Sender", "host", cfg.SMTPHost)
	} else {
		sender = &service.MockSender{}
		logger.Info("Using Mock Sender")
	}

	// Setup Handler
	h := handlers.NewHandler(repo, sender)

	// Setup Router
	r := chi.NewRouter()
	r.Use(middleware.RequestLogger(logger))
	r.Use(chiMiddleware.Recoverer)

	// Basic CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-API-Key"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Post("/register", h.Register)
	r.Post("/login", h.Login)
	r.Post("/trigger/{token}", h.TriggerEmail)

	r.Group(func(r chi.Router) {
		r.Use(middleware.JWTAuth)
		r.Post("/api-keys", h.GenerateAPIKey)
		r.Get("/api-keys", h.ListAPIKeys)

		r.Get("/templates", h.ListTemplates)
		r.Post("/templates", h.CreateTemplate)
		r.Get("/templates/{id}", h.GetTemplate)
		r.Put("/templates/{id}", h.UpdateTemplate)
		r.Delete("/templates/{id}", h.DeleteTemplate)

		r.Get("/logs", h.GetEmailLogs)
	})

	r.Group(func(r chi.Router) {
		r.Use(middleware.APIKeyAuth(repo))
		r.Post("/send", h.SendEmail)
	})

	logger.Info("Starting server", "port", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		logger.Error("Server failed", "error", err)
		os.Exit(1)
	}
}
