package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

const PORT = "8086"

func prettyPrintJson(d any) string {
	b, err := json.MarshalIndent(d, "", "    ")
	if err != nil {
		log.Println("error while pretty printing json")
		return ""
	}
	return string(b)
}

func main() {
	httpClient := &http.Client{}

	// waitgroup to manage goroutines(green threads)
	wg := &sync.WaitGroup{}

	// create 5 clients and make 3 requests per client to the server concurrently, using green threads
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		// goroutines(green threads)
		go func() {
			k := i
			defer wg.Done()

			// for each client make 3 requests
			for j := 0; j < 3; j++ {
				MakeClientRequest(httpClient, k)
			}
		}()
	}

	// wait for the goroutines to finish
	wg.Wait()
}

// MakeClientRequest makes a POST request to the server and prints the response
func MakeClientRequest(httpClient *http.Client, clientID int) {

	// json data that we are sending to the server
	p := map[string]string{"client_id": strconv.Itoa(clientID), "message": "ping"}

	// json serialization
	b, err := json.Marshal(p)
	if err != nil {
		log.Printf("client_id: %v error while serializing payload : %v\n", clientID, err)
		return
	}

	// create a new http request
	req, err := http.NewRequest(http.MethodPost, "http://"+os.Getenv("SERVER_NAME")+":"+PORT, bytes.NewBuffer(b))
	if err != nil {
		log.Printf("client_id: %v error while creating new request :: %v\n", clientID, err)
		return
	}

	req.Header.Set("client_id", strconv.Itoa(clientID))

	// make the http request
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Printf("client_id: %v error while making request :: %v\n", clientID, err)
		return
	}

	// close the body after it is done reading
	defer resp.Body.Close()

	m := map[string]string{}

	// deserialize the response
	err = json.NewDecoder(resp.Body).Decode(&m)
	if err != nil {
		log.Printf("client_id: %v error while reading body : %v\n", clientID, err)
		return
	}

	// print the response
	fmt.Println(prettyPrintJson(m))
}
