package handler

import (
	"encoding/json"
	"fmt"
	"github.com/AtaullinShamil/L0/pkg/db"
	"github.com/go-chi/chi"
	"github.com/yuin/goldmark/renderer"
	"io"
	"sync"
)

type Handler struct {
	db       *db.Database
	renderer renderer.Renderer
	cache    *sync.Map

	url string
}

func NewHandler(db *db.Database, url string, cache *sync.Map) *Handler {
	return &Handler{
		db:    db,
		url:   fmt.Sprintf("%s/api/v1", url),
		cache: cache,
	}
}

func (h *Handler) Router() chi.Router {
	router := chi.NewRouter()

	router.Post("/check", h.GetInfo)
	router.Get("/wb", h.GetHTML)

	return router
}

func (h *Handler) parseBody(from io.ReadCloser, to interface{}) error {
	body, err := io.ReadAll(from)
	if err != nil || len(body) == 0 {
		return fmt.Errorf("empty body")
	}
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, to)
	if err != nil {
		return err
	}
	return nil
}
