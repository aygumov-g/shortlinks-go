package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/aygumov-g/shortlinks-go/application/src/internal/app/web/handlers/form"
	"github.com/aygumov-g/shortlinks-go/application/src/internal/app/web/handlers/home"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("../.env")
	mux := http.NewServeMux()
	mux.HandleFunc("/", home.Handler())
	mux.HandleFunc("/form", form.Handler())
	if err := http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), mux); err != nil {
		panic(fmt.Sprintf("Server is not running! Error=%s", err))
	}
}
