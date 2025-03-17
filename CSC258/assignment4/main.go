package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Simulate memory allocation (allocating 100MB on each request)
	size := 100 * 1024 * 1024 // 100 MB
	allocatedMemory := make([]byte, size)
	for i := 0; i < size; i++ {
		allocatedMemory[i] = byte(i % 256) // Fill with some data
	}

	//delay to further simulate load
	time.Sleep(1 * time.Second)

	log.Println("request received")
	// Respond to the client
	fmt.Fprintf(w, "Memory load simulated. Allocated %d MB.\n", size/(1024*1024))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
