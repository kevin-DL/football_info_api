package main

import (
	"football_api/ent"
	"football_api/general"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	client, err := ent.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_URL"))

	if err != nil {
		log.Fatalf("failed opening connection to %s: %v", os.Getenv("DB_DRIVER"), err)
	}
	defer client.Close()

	server := CreateServer(client)
	server.SetSuperTokens()
	server.SetupMiddlewares()
	server.SetupHandlers()

	log.Println("Running on port", port)
	err = http.ListenAndServe(":"+port, server.Router)
	if err != nil {
		log.Fatalf("Failed to run server: %s", err)
	}
}

func CreateServer(client *ent.Client) *general.Server {
	server := &general.Server{
		Router: chi.NewRouter(),
		Client: client,
	}
	server.Router.Get("/health", general.HandleHealth)
	return server
}
