package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {

	message := "Pass a name in the query string for a personalized response. \n"
	name := r.URL.Query().Get("name")

	if name != "" {
		message = fmt.Sprintf("Hello, %s!\n", name)
	}

	fmt.Fprint(w, message)
}

type InvokeRequest struct {
	Data     map[string]json.RawMessage
	Metadata map[string]interface{}
}

type InvokeResponse struct {
	Outputs     map[string]interface{}
	Logs        []string
	ReturnValue interface{}
}

func queueTriggerHandler(w http.ResponseWriter, r *http.Request) {

	var invokeRequest InvokeRequest

	d := json.NewDecoder(r.Body)
	d.Decode(&invokeRequest)

	item := invokeRequest.Data["myQueueItem"]

	outputs := map[string]interface{}{"": ""}

	invokeResponse := InvokeResponse{outputs, nil, ""}

	invokeResponse.Logs = append(invokeResponse.Logs, "hello from queue trigger")
	invokeResponse.Logs = append(invokeResponse.Logs, string(item))

	responseJson, _ := json.Marshal(invokeResponse)
	fmt.Println(string(responseJson))
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJson)
}

func main() {
	listenAddr := ":8080"

	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}

	http.HandleFunc("/api/hello", helloHandler)
	http.HandleFunc("/queueTrigger", queueTriggerHandler)
	log.Printf("About to listen on %s. Go to http://127.0.0.1%s", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
