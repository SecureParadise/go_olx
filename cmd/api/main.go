package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/SecureParadise/olx_monolith/internal/config"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println("starting olx server .... ")
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

	server := &http.Server{
		// Addr: ":8090",
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}

	// err := http.ListenAndServe(":8090", mux)
	log.Printf("server is listening on %s", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
