package handler

import (
	"net/http"
)

func (h *Handler) GetInfo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
}
