package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	// mux = router
	// router itself is atype of handler in golang
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		// http.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		// strictly follow chronollogy and order
		w.Write([]byte(`{"status":"ok"}`))
	})

	server := http.Server{
		Addr:         ":8090",
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}

	// err := http.ListenAndServe(":8090", mux)
	err := server.ListenAndServe()

	if err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
