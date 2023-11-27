package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func (h *Handler) GetHTML(w http.ResponseWriter, r *http.Request) {
	htmlFile, err := os.Open("/Users/shamil/Desktop/L0/public/order_check.html")

	if err != nil {
		log.Fatal(err)
	}
	defer htmlFile.Close()

	htmlInfo, _ := htmlFile.Stat()
	size := htmlInfo.Size()
	bytes := make([]byte, size)
	htmlFile.Read(bytes)
	fmt.Fprintf(w, string(bytes))
}
