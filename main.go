package main

import (
	"dictionary/dictionary"
	"fmt"
)

func main() {
	filepath := "dictionary/dict.json"
	d := dictionary.NewDictionary(filepath)

	d.Add("francais", "french")
	d.Add("anglais", "english")
	d.Add("espagnol", "spanish")
	fmt.Println(d)
	d.Add("italien", "italian")
	v1, _ := d.Get("anglais")

	fmt.Println("v1:", v1)

	d.Remove("anglais")

	list, err := d.List()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("List:", list)
	}

	
	
}
