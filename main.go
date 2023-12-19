package main

import (
	"fmt"
)



func main() {
	m := make(map[string]string)
	m["francais"] = "french"
	m["anglais"] = "english"
    m["espagnol"] = "spanish"
	

	v1 := get(m, "anglais")

	fmt.Println("map:", m)
	fmt.Println("v1:", v1)
	
	remove(m, "anglais")
	fmt.Println("map:", m)
    fmt.Println("List:", list(m))
	

	
}
func get(m map[string]string, key string) string {
	return m[key]
}

func remove(m map[string]string, key string) {
	delete(m, key)
}

func list(m map[string]string) []string {
    result := make([]string, 0, len(m))
    for key, value := range m {
        result = append(result, key+": "+value+",")
    }
    return result
}
