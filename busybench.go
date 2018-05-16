package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"cloud.google.com/go/profiler"
)

const size = 1024 * 1024

// busywork continuously generates 1MiB of random data and compresses it
// throwing away the result.
func busywork() {
	for {
		busyworkOnce()
	}
}

func busyworkOnce() {
	data := make([]byte, size)
	rand.Read(data)

	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write(data); err != nil {
		log.Fatalf("Failed to write to gzip stream", err)
	}
	if err := gz.Flush(); err != nil {
		log.Fatalf("Failed to flush to gzip stream", err)
	}
	if err := gz.Close(); err != nil {
		log.Fatalf("Failed to close gzip stream", err)
	}
	// Throw away the result.
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "called %s", r.URL.Path[1:])
}

func main() {
	err := profiler.Start(
		profiler.Config{
			Service:        "busybench",
			ServiceVersion: "1.0.0",
			DebugLogging:   true,
		})
	if err != nil {
		log.Fatalf("Failed to start the profiler: %v", err)
	}

	go busywork()

	httpAddr := ":8080"
	log.Printf("Starting an HTTP server on %s", httpAddr)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(httpAddr, nil))
}
