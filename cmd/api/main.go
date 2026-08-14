package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/SecureParadise/olx_monolith/internal/config"
	"github.com/SecureParadise/olx_monolith/internal/handlers"
)

// the major funcanality of main.go is to wireup every funcanality
func main() {
	cfg := config.MustLoad()
	fmt.Println("starting olx server .... ")
	// mux = router
	// router itself is atype of handler in golang
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handlers.Health)

	server := &http.Server{
		// Addr: ":" + os.Getenv("PORT"),
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
