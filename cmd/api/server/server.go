package server

import (
	"github.com/artcurty/go-proxy-make/internal"
	"github.com/artcurty/go-proxy-make/pkg"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func StartServer() {
	port := pkg.GetEnv("SERVER_PORT", ":8081")
	r := mux.NewRouter()

	internal.RegisterRoutes(r)

	log.Printf("Starting server on port %s\n", port)
	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Fatalf("Server failed to start: %v\n", err)
	}
}
