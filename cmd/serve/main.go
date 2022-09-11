package main

import (
	"io/ioutil"
	"lambdacommerce/internal/app"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	response, _ := app.HandleRequest(app.Request{
		Body:   string(body),
		Path:   r.URL.Path,
		Method: r.Method,
	})

	w.WriteHeader(response.StatusCode)
	w.Header().Set("Content-Type", "application/text")
	w.Write([]byte(response.Body))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	})

	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}