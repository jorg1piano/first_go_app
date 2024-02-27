package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Starting server")

	godotenv.Load(".env")

	portstring := os.Getenv("PORT")
	if portstring == "" {
		log.Fatal("Port not set")
	}

	router := chi.NewRouter()

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portstring,
	}

	log.Printf("Server starting on port %s", portstring)
	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("PORT: ", portstring)
}
