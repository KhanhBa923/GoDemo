package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
func (app *application) hello(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"message": "Hello, World!",
	}

	// Thiết lập header để trình duyệt hiểu là JSON
	w.Header().Set("Content-Type", "application/json")

	// Mã hóa JSON và gửi về client
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
