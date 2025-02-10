package main

import (
	"encoding/json"
	"fetch-assessment/api"
	"net/http"
)

// register custom routes for the server (that aren't provided in the spec)
// this is just to provide swagger for local testing and a heath check
func registerCustomRoutes(mux *http.ServeMux) {
	// register swagger to testing
	mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		spec, err := api.GetSwagger()
		if err != nil {
			http.Error(w, "Failed to load Swagger spec", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(spec)
	})

	// register health check
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
}
