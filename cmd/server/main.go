package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/omarmendozaaa/auth-master/internal/config"
	"github.com/omarmendozaaa/auth-master/internal/models"
)

func main() {
	config.LoadConfig()
	config.ConnectDatabase()

	err := config.DB.AutoMigrate(
		&models.Project{},
		&models.User{},
		&models.Role{},
		&models.UserProject{},
	)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	fmt.Println("Database migrated successfully")

	port := config.Cfg.Port

	fmt.Printf("🚀 Auth Master running on port %d in %s mode\n", port, config.Cfg.Env)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
