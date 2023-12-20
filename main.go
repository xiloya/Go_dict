// main.go

package main

import (
	"dictionary/dictionary"
	"log"
	"net/http"
)

func main() {
	filepath := "dictionary/dict.json"
	d := dictionary.NewDictionary(filepath)

	http.HandleFunc("/add", d.AddHandler)
	http.HandleFunc("/get", d.GetHandler)
	http.HandleFunc("/remove", d.RemoveHandler)
	http.HandleFunc("/list", d.ListHandler)

	port := "8080"
	log.Printf("Server listening on port %s...\n", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("Error starting server:", err)
	}
	
}
