package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {

	message := "Pass a name in the query string for a personalized response. \n"
	name := r.URL.Query().Get("name")

	if name != "" {
		message = fmt.Sprintf("Hello, %s!\n", name)
	}

	fmt.Fprint(w, message)
}

func main() {
	listenAddr := ":8080"
	http.HandleFunc("/api/hello", helloHandler)
	log.Printf("About to listen on %s. Go to http://127.0.0.1%s", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
