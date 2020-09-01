package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// WebAPIResult Struct To Serve as the Result Of A WebService Request
type WebAPIResult struct {
	Status int    `json:"Status"`
	Body   string `json:"Body"`
}

// WebAPIResult2 struct basically a person but lets go with it
type WebAPIResult2 struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Height    int    `json:"Height"`
	JobTitle  string `json:"JobTitle"`
}

func main() {
	handleRequests()
}

func handleRequests() {
	http.HandleFunc("/home", homePage) // Path and handler method, handler must take an http resonse writer and the request object http://localhost:10000/home
	http.HandleFunc("/Person", secondPage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	// Create result object
	var result = WebAPIResult{Status: 200, Body: "Pretend This Response has data that is valuable to the consumer!"}

	// Encode result object as json and write it to w
	json.NewEncoder(w).Encode(result)

	// Log that the endpoint was hit to the console
	fmt.Println("Endpoint Hit: at /home on port 10000")
}

func secondPage(w http.ResponseWriter, r *http.Request) {
	// Create result object
	var result = WebAPIResult2{FirstName: "Joe", LastName: "Franks", Height: 6, JobTitle: "Software Engineer"}

	// Encode result object as json and write it to w
	json.NewEncoder(w).Encode(result)

	// Log that the endpoint was hit to the console
	fmt.Println("Endpoint Hit: at /person on port 10000")
}
