package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/omarmendozaaa/auth-master/internal/config"
)

func main() {
	config.LoadConfig()
	config.ConnectDatabase()
	port := config.Cfg.Port

	fmt.Printf("🚀 Auth Master running on port %d in %s mode\n", port, config.Cfg.Env)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
