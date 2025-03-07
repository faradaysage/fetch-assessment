package main

import (
	"context"
	"fetch-assessment/api"
	"fetch-assessment/repository"
	"fetch-assessment/server"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// create an in-memory repository
	repository := repository.NewMemoryRepository()

	// create a new server instance using the in-memory repository
	serverLogic := server.NewServer(repository)

	// generate a strict handler from the server implementation with custom error handling
	strictHandler := api.NewStrictHandlerWithOptions(serverLogic, nil, api.StrictHTTPServerOptions{
		RequestErrorHandlerFunc:  customErrorHandler,
		ResponseErrorHandlerFunc: customErrorHandler,
	})

	// create a new multiplexer router
	mux := http.NewServeMux()

	// register custom routes
	registerCustomRoutes(mux)

	// register the handler with the router
	api.HandlerFromMux(strictHandler, mux)

	// wrap the mux with middleware
	handlerWithMiddleware := addCustomMiddleware(mux)

	// define the port to listen on
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// create an http server
	server := &http.Server{
		Addr:    ":" + port,
		Handler: handlerWithMiddleware,
	}

	// create a channel to listen for OS interrupt or termination signals.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// run server asynchronously so we're not blocking
	// for this singular use case, it's not a big deal, but we could make this one of many services we run fron this process
	go func() {
		log.Printf("Server listening on port %s (Press Ctrl+C to stop)", port)
		// ListenAndServe will return on error
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
		log.Println("Server stopped.")
	}()

	// block until we receive a signal on the quit channel
	<-quit

	log.Println("Shutdown signal received, shutting down gracefully...")

	// give the server 5 seconds to shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt the graceful shutdown by closing the listener and completing ongoing requests.
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server encountered error while shutting down: %v", err)
	} else {
		log.Println("Server shutdown completed gracefully")
	}
}
