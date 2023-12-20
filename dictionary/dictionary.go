package dictionary

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Dictionary struct {
	filePath string
}
type Data struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func NewDictionary(filePath string) Dictionary {
	return Dictionary{
		filePath: filePath,
	}
}

func (d *Dictionary) Get(key string) (string, error) {
	data, err := os.ReadFile(d.filePath)
	if err != nil {
		return "", err
	}

	var jsonData map[string]string
	if err := json.Unmarshal(data, &jsonData); err != nil {
		return "", err
	}

	value, exists := jsonData[key]
	if !exists {
		return "", errors.New("key not found")
	}

	return value, nil
}

func (d *Dictionary) Remove(key string) error {
	file, err := os.ReadFile(d.filePath)
	if err != nil {
		return err
	}

	var jsonData map[string]string
	if err := json.Unmarshal(file, &jsonData); err != nil {
		return err
	}

	// Check if the key exists
	if _, exists := jsonData[key]; !exists {
		return errors.New("key not found")
	}

	delete(jsonData, key)

	updatedData, err := json.MarshalIndent(jsonData, "", "    ")
	if err != nil {
		return err
	}

	err = os.WriteFile(d.filePath, updatedData, os.ModePerm)
	if err != nil {
		return err
	}

	fmt.Println("removed successfully!")
	return nil
}

func (d Dictionary) Add(key, value string) error {
	dataMap := make(map[string]string)
	if existingData, err := os.ReadFile(d.filePath); err == nil && len(existingData) > 0 {
		if err := json.Unmarshal(existingData, &dataMap); err != nil {
			return err
		}
	}
	dataMap[key] = value

	updatedData, err := json.MarshalIndent(dataMap, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(d.filePath, updatedData, 0644)
	if err != nil {
		return err
	}

	fmt.Println(" added successfully!")
	return nil
}
func (d *Dictionary) List() ([]string, error) {
	data, err := os.ReadFile(d.filePath)
	if err != nil {
		return nil, err
	}

	var jsonData map[string]string
	if err := json.Unmarshal(data, &jsonData); err != nil {
		return nil, err
	}

	var result []string
	for key, value := range jsonData {
		result = append(result, fmt.Sprintf("%s: %s", key, value))
	}

	return result, nil
}