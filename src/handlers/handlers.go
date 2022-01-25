package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)


func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Pig!!")
}

func PostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var metadata MetaData
	if err := decoder.Decode(&metadata); err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	fmt.Fprintf(w, "Request Payload: %v\n", metadata)
}

func CreateRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Everything seems to be fine!")
}

type MetaData interface{}
