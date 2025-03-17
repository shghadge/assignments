package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"os"
	"sync"
	"time"
)

const PORT = "8086"

type Payload struct {
	ClientID string `json:"client_id"`
	Message  string `json:"message"`
}

// RecoverMiddleware will catch any panic and log it without crashing the server
func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// defer a function to handle the panic
		defer func() {
			if r := recover(); r != nil {
				// handle the panic and return a proper response
				log.Printf("Recovered from panic: %v", r)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		// call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}

// LoggingMiddleware wraps a http handler and logs details about the incoming request
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// log the incoming request
		log.Printf("Request: %s %s, Client ID: %s", r.Method, r.URL.Path, r.Header.Get("client_id"))
		// call the next handler in the middelware chain
		next.ServeHTTP(w, r)
	})
}

// handler handles the incoming http request and returns additional details to the client
func (s *server) handler(w http.ResponseWriter, r *http.Request) {
	receivedTime := time.Now()
	p := Payload{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Println("error while decoding body", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"type": "error", "msg": "error while reading body"})
		return
	}

	// Write client ID and access time to log file with lock
	s.mutex.Lock()
	log.Printf("ClientID: %s acquired the lock", p.ClientID)

	// defer calls get executed at the end of the function
	defer func() {
		log.Printf("ClientID: %s releasing the lock", p.ClientID)
		s.mutex.Unlock() // lock will be realeased before the function return
	}()

	// open the log file. create and open if it doesn't exist
	logFile, err := os.OpenFile("log_file", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("error opening log file:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"type": "error", "msg": "internal server error"})
		return
	}
	defer logFile.Close()

	// write to the log file when the client request is received
	logEntry := fmt.Sprintf("ClientID: %s, Accessed at: %s\n", p.ClientID, receivedTime.UTC().String())
	_, err = logFile.WriteString(logEntry)
	if err != nil {
		log.Println("error writing to log file:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"type": "error", "msg": "error writing log"})
		return
	}

	// sleep for a random time to simulate processing
	randomNum := rand.IntN(4-1) + 1
	time.Sleep(time.Duration(randomNum) * time.Second)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"type": "data", "client_id": p.ClientID, "received_at": receivedTime.UTC().String(), "processed_time": time.Now().UTC().String(), "message": "pong"})
}

type server struct {
	mutex sync.RWMutex
}

func main() {
	s := server{
		mutex: sync.RWMutex{},
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", s.handler)

	// add the logging and panic recover middleware
	configuredMux := RecoverMiddleware(LoggingMiddleware(mux))

	log.Println("Started Listening on Port", PORT)
	// serve the router
	log.Fatal(http.ListenAndServe(":"+PORT, configuredMux))
}
