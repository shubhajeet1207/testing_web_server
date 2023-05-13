package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Create a new web server that writes the prompt to all requests.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //func(w http.ResponseWriter, r *http.Request) is function signature which has two params: write responses(http.ResponseWriter) and represent http request(http.Request)
		fmt.Fprintf(w, "Hello, world! This is a test server.\n")
	})

	// Listen for requests for port 8080.
	log.Fatal(http.ListenAndServe(":8080", nil)) //log.Fatal() prints out the error message and exits the program and "http.ListenAndServe"  listens the requets on the specified port(8080) and handles the requests.

	// Listen for requests on all ports.
	//  for _, port := range []int{80, 443} {
	//     go func(port int) {
	//         log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
	//     }(port)
	// }
}
