package main

import (
	"fetch-assessment/server"
	"log"
	"net/http"
)

func customErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Request error: %v", err)
	response := server.CustomPostReceiptsProcess400Response{}
	if err := response.VisitPostReceiptsProcessResponse(w); err != nil {
		log.Printf("Failed to write 400 response: %v", err)
	}
}
