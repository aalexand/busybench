package main

import (
	"bytes"
	"compress/gzip"
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

func main() {
	err := profiler.Start(
		&profiler.Config{
			Target:       "busybench",
			DebugLogging: true,
		})
	if err != nil {
		log.Fatalf("Failed to start the profiler: %v", err)
	}

	go busywork()

	httpAddr := "localhost:8080"
	log.Printf("Starting an HTTP server on %s", httpAddr)
	log.Fatal(http.ListenAndServe(httpAddr, nil))
}