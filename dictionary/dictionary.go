package dictionary

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type KeyValuePair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Dictionary struct {
	filePath string
}

func NewDictionary(filePath string) Dictionary {
	return Dictionary{
		filePath: filePath,
	}
}

func (d *Dictionary) AddHandler(w http.ResponseWriter, r *http.Request) {
	var entry KeyValuePair
	if err := json.NewDecoder(r.Body).Decode(&entry); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	dataMap := make(map[string]string)
	file, _ := os.ReadFile(d.filePath)
	_ = json.Unmarshal(file, &dataMap)

	dataMap[entry.Key] = entry.Value

	updatedData, _ := json.MarshalIndent(dataMap, "", "  ")
	_ = os.WriteFile(d.filePath, updatedData, 0644)

	w.WriteHeader(http.StatusCreated)
}

func (d *Dictionary) GetHandler(w http.ResponseWriter, r *http.Request) {
	var key string

	queryKey := r.URL.Query().Get("key")
	if queryKey == "" {
		var requestBody struct {
			Key string `json:"key"`
		}

		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		key = requestBody.Key
	} else {
		key = queryKey
	}

	if key == "" {
		http.Error(w, "Missing 'key' parameter", http.StatusBadRequest)
		return
	}

	dataMap := make(map[string]string)
	file, _ := os.ReadFile(d.filePath)
	_ = json.Unmarshal(file, &dataMap)

	value, exists := dataMap[key]
	if !exists {
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	}

	response := KeyValuePair{Key: key, Value: value}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func (d *Dictionary) RemoveHandler(w http.ResponseWriter, r *http.Request) {
	var key string

	queryKey := r.URL.Query().Get("key")
	if queryKey == "" {
		var requestBody struct {
			Key string `json:"key"`
		}

		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		key = requestBody.Key
	} else {
		key = queryKey
	}

	if key == "" {
		http.Error(w, "Missing 'key' parameter", http.StatusBadRequest)
		return
	}

	dataMap := make(map[string]string)
	file, _ := os.ReadFile(d.filePath)
	_ = json.Unmarshal(file, &dataMap)

	if _, exists := dataMap[key]; !exists {
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	}

	delete(dataMap, key)

	updatedData, _ := json.MarshalIndent(dataMap, "", "    ")
	_ = os.WriteFile(d.filePath, updatedData, os.ModePerm)

	w.WriteHeader(http.StatusNoContent)
}

func (d *Dictionary) ListHandler(w http.ResponseWriter, r *http.Request) {
	dataMap := make(map[string]string)
	file, _ := os.ReadFile(d.filePath)
	_ = json.Unmarshal(file, &dataMap)

	var result []string
	for key, value := range dataMap {
		result = append(result, fmt.Sprintf("%s: %s", key, value))
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(result)
}
