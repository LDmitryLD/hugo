package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"projects/LDmitryLD/hugoproxy/proxy/internal/infrastructure/router"
	"projects/LDmitryLD/hugoproxy/proxy/internal/modules"
	"syscall"
	"time"
)

func main() {
	r := router.NewRouter(modules.NewControllers())

	server := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("Starting server...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %s", err.Error())
		}
	}()

	<-sigChan

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("Server stopped gracefully")
}
