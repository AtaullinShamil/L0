package handler

import (
	"fmt"
	"net/http"
	"os"
)

func (h *Handler) GetHTML(w http.ResponseWriter, r *http.Request) {
	htmlFile, err := os.Open(h.path + "/public/order_check.html")

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer htmlFile.Close()

	htmlInfo, _ := htmlFile.Stat()
	size := htmlInfo.Size()
	bytes := make([]byte, size)
	htmlFile.Read(bytes)
	fmt.Fprintf(w, string(bytes))
}
