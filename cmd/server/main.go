package main

import (
	"context"
	"go_tutor/internal/handler"
	"go_tutor/internal/repository"
	"go_tutor/internal/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)

	r := gin.Default()

	productRepo := repository.NewMemoryRepository()
	productService := service.NewProductService(productRepo)
	handler.NewProductHandler(r, productService)

	srv := &http.Server{
		Addr: ":8080",
		Handler: r,
	}

	go func() {
		log.Printf("Server is running on port %s\n", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to start server: %v\n", err)
		}
	}()

	quit:= make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server..")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err!= nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting.")
}