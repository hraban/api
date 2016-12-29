package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func yesliekttly(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/plain; charset=utf-8; brains=tasty")
	fmt.Fprintf(w, "yes this program is basically the best in the world.")
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Printf("Listening on %s", port)
	http.HandleFunc("/isthisprogramawesome", yesliekttly)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
