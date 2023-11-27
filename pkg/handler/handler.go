package handler

import (
	"fmt"
	"github.com/AtaullinShamil/L0/pkg/db"
	"github.com/go-chi/chi"
	"github.com/yuin/goldmark/renderer"
)

type Handler struct {
	db       *db.Database
	renderer renderer.Renderer

	url string
}

func NewHandler(db *db.Database, url string) *Handler {
	return &Handler{
		db:  db,
		url: fmt.Sprintf("%s/api/v1", url),
	}
}

func (h *Handler) Router() chi.Router {
	router := chi.NewRouter()

	router.Get("/wb/check", h.GetInfo)

	return router
}
