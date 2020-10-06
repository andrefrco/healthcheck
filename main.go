package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "healthcheck_processed_ops_total",
		Help: "The total number of processed events",
	})
)

func main() {
	recordMetrics()

	var listen_port string = ":" + getEnv("LISTEN_PORT", "8000")
	router := mux.NewRouter()
	router.HandleFunc("/healthcheck", HealthCheck).Methods("GET")
	router.Handle("/metrics", promhttp.Handler())
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
