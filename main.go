package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type HealthResponse struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(HealthResponse{Status: "ok", Version: appVersion()})
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from go-web-app! Version: " + appVersion()))
}

func appVersion() string {
	if v := os.Getenv("APP_VERSION"); v != "" {
		return v
	}
	return "dev"
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/healthz", healthHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("listening on :%s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}