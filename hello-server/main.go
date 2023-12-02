package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		const chunkSize = 50 * 1024 * 1024
		var memoryChunk []byte

		for {
			newData := make([]byte, chunkSize)
			for i := 0; i < chunkSize; i++ {
				newData[i] = 0
			}

			memoryChunk = append(memoryChunk, newData...)

			// 少しの遅延を追加してCPUの過剰使用を避ける
			time.Sleep(10 * time.Millisecond)
		}
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/healthz" {
			http.NotFound(w, r)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK")
		log.Printf("Health Status OK")
	})

	http.Handle("/metrics", promhttp.Handler())

	log.Printf("Starting server on port %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
