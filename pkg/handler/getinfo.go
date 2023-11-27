package handler

import (
	"encoding/json"
	"github.com/AtaullinShamil/L0/pkg/db"
	"net/http"
)

func (h *Handler) GetInfo(w http.ResponseWriter, r *http.Request) {
	var request Request
	err := h.parseBody(r.Body, &request)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	jsonData, ok := h.cache.Load(request.Uid)
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	var s db.Order
	err = json.Unmarshal(jsonData.([]byte), &s)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(s); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}
