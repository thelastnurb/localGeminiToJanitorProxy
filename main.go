package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"

	"proxy/handler"
	"proxy/middleware"
)

func main() {
	_ = godotenv.Load() // Игнорируем ошибку, если файла нет, используем env vars системы

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if os.Getenv("GEMINI_API_KEY") == "" {
		log.Fatal("GEMINI_API_KEY is not set in environment variables")
	}

	mux := http.NewServeMux()

	// Endpoints
	mux.HandleFunc("GET /health", handler.HealthHandler)
	mux.HandleFunc("GET /v1/models", handler.ModelsHandler)
	mux.HandleFunc("POST /v1/chat/completions", handler.CompletionsHandler)

	// Middlewares: Logger -> CORS -> Auth
	handler := middleware.Logger(middleware.CORS(middleware.Auth(mux)))

	server := &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}

	// Graceful shutdown
	go func() {
		log.Printf("Starting proxy server on port %s", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}