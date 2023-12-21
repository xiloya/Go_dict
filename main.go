package main

import (
	"dictionary/dictionary"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sync"
)

func main() {
	filepath := "dictionary/dict.json"
	d := dictionary.NewDictionary(filepath)

	http.HandleFunc("/add", d.AddHandler)
	http.HandleFunc("/get", d.GetHandler)
	http.HandleFunc("/remove", d.RemoveHandler)
	http.HandleFunc("/list", d.ListHandler)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		port := "8080"
		log.Printf("Server listening on port %s...\n", port)
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			log.Fatal("Error starting server:", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			entry := <-dictionary.AddChannel
			dataMap := make(map[string]string)
			file, _ := os.ReadFile(filepath)
			_ = json.Unmarshal(file, &dataMap)

			dataMap[entry.Key] = entry.Value

			updatedData, _ := json.MarshalIndent(dataMap, "", "  ")
			_ = os.WriteFile(filepath, updatedData, 0644)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			key := <-dictionary.RemoveChannel
			dataMap := make(map[string]string)
			file, _ := os.ReadFile(filepath)
			_ = json.Unmarshal(file, &dataMap)

			if _, exists := dataMap[key]; exists {
				delete(dataMap, key)

				updatedData, _ := json.MarshalIndent(dataMap, "", "    ")
				_ = os.WriteFile(filepath, updatedData, 0644)
			}
		}
	}()

	wg.Wait()

	close(dictionary.AddChannel)
	close(dictionary.RemoveChannel)
}
