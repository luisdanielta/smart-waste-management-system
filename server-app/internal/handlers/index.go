package handlers

import (
	"net/http"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Not found", http.StatusNotFound)
		return
	}
}
