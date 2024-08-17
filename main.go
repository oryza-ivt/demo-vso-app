package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// Main handler for the HTTP server
func handler(w http.ResponseWriter, r *http.Request) {
	// Pod name from the Downward API (injected via environment variable)
	podName := os.Getenv("POD_NAME")
	if podName == "" {
		podName = "Pod name not found"
	}

	// Secret from environment variables (injected via Kubernetes Secrets)
	secretPassword := os.Getenv("SECRET_PASSWORD")
	if secretPassword == "" {
		secretPassword = "Secret password not found"
	}

	// Secret from environment variables (injected via Kubernetes Secrets)
	secretUsername := os.Getenv("SECRET_USERNAME")
	if secretUsername == "" {
		secretUsername = "Secret username not found"
	}

	fmt.Fprintf(w, "Pod Name: %s\n", podName)
	fmt.Fprintf(w, "Secret - username: %s\n", secretUsername)
	fmt.Fprintf(w, "Secret - password: %s\n", secretPassword)
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Version: %s\n", "0.2")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/version", versionHandler)
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
