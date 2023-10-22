package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/coby-amar/go_learning/main/config"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	apiConf := config.Setup()
	if apiConf == nil {
		slog.Error("CreateApiConfig returned nil")
		return
	}
	router := setupRouter(apiConf)
	initServer(router)
}

func initServer(router *chi.Mux) {
	port := os.Getenv("PORT")
	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}
	slog.Info(fmt.Sprintln("Staring server on port:", port))
	err := server.ListenAndServe()
	if err != nil {
		slog.Error("Fatal server couldnt start port:")
	}
}
