package main

import (
	"encoding/json"
	"log"
	"math/rand/v2"
	"net/http"
	"time"
)

const PORT = "8086"

type Payload struct {
	ClientID string `json:"client_id"`
	Message  string `json:"message"`
}

// A simple handler that triggers a panic
func panicHandler(w http.ResponseWriter, r *http.Request) {
	// This will cause a panic
	panic("Something went wrong!")
}

// handler handles the incoming http request and returns additional details to the client
func handler(w http.ResponseWriter, r *http.Request) {
	receivedTime := time.Now()
	p := Payload{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"type": "error", "msg": "error while reading body"})
		return
	}

	// sleep for a random time to simulate processing
	randomNum := rand.IntN(4-1) + 1
	time.Sleep(time.Duration(randomNum) * time.Second)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"type": "data", "client_id": p.ClientID, "received_at": receivedTime.UTC().String(), "processed_time": time.Now().UTC().String(), "message": "pong"})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	// add the logging middleware
	configuredMux := RecoverMiddleware(LoggingMiddleware(mux))

	log.Println("Started Listening on Port", PORT)
	// serve the router
	log.Fatal(http.ListenAndServe(":"+PORT, configuredMux))
}
