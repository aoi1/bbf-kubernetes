package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const defaultMessage = "Hello, world! Let's learn Kubernetes!"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	message, err := readMessageFromConfig()
	if err != nil {
		log.Fatalf("Failed to read message from config: %v", err)
		message = defaultMessage
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, message)
	})

	log.Printf("Starting server on port %s\n", port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func readMessageFromConfig() (string, error) {
	configDir := "/etc/config"
	configPath := fmt.Sprintf("%s/myconfig.txt", configDir)

	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return defaultMessage, nil
	}

	return string(data), nil
}
