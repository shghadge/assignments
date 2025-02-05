package main

import (
	"log"
	"net/http"
)

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
