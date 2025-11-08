package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8080
	fmt.Printf("🚀 Auth Master running on port %d\n", port)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
