package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Printf("Listening on %s", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
