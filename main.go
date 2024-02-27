package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("hello")

	godotenv.Load(".env")

	portstring := os.Getenv("PORT")
	if portstring == "" {
		log.Fatal("Port not set")
	}

	fmt.Println("PORT: ", portstring)
}
