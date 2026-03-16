package main

import (
	"log"
	"net/http"
	"os"

	httpadapter "github.com/boosun13/ArcNote/backend/internal/adapter/http"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	addr := serverAddr()
	log.Printf("ArcNote API starting on http://localhost%s", addr)
	return http.ListenAndServe(addr, httpadapter.NewServer())
}

func serverAddr() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return ":" + port
}
