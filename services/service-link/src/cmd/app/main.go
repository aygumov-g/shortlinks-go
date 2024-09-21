package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aygumov-g/shortlinks-go/services/service-link/src/internal/app/storage"
	"github.com/aygumov-g/shortlinks-go/services/service-link/src/internal/app/web/home"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	godotenv.Load("../.env")

	storageLink, err := storage.NewStorage()
	if err != nil {
		panic(err)
	}

	serviceLink := home.NewService(storageLink)

	router := mux.NewRouter()
	router.HandleFunc("/", serviceLink.CreateLink).Methods("POST")
	router.HandleFunc("/{link_addr_in}", serviceLink.LinkSearch).Methods("GET")

	c := cors.New(cors.Options{
		//AllowedOrigins:   []string{"http://localhost:3333"},
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	log.Printf("[Link]: Server is running on port: %s\n", os.Getenv("SERVER_PORT"))
	log.Fatal(
		http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")), handler),
	)
}
