package handlers

import (
	"egogo/internal/repository"
	"egogo/internal/service"
)

type Handler struct {
	Repo   repository.Repository
	Sender service.EmailSender
}

func NewHandler(repo repository.Repository, sender service.EmailSender) *Handler {
	return &Handler{
		Repo:   repo,
		Sender: sender,
	}
}
