package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	var listen_port string = ":" + getEnv("LISTEN_PORT", "8000")
	router := mux.NewRouter()
	router.HandleFunc("/healthcheck", HealthCheck).Methods("GET")
	log.Fatal(http.ListenAndServe(listen_port, router))
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "UP")
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
