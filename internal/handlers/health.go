package handlers

import "net/http"

func Health(w http.ResponseWriter, r *http.Request) {
	// http.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// strictly follow chronollogy and order
	w.Write([]byte(`{"status":"okey rey"}`))
}
