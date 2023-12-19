package main

import (
	"dictionary/dictionary"
	"fmt"
)

func main() {
	d := dictionary.NewDictionary()

	d.Add("francais", "french")
	d.Add("anglais", "english")
	d.Add("espagnol", "spanish")
	fmt.Println(d)
	d.Add("italien", "italian")
	v1 := d.Get("anglais")

	
	fmt.Println("v1:", v1)

	d.Remove("anglais")
	
	fmt.Println("List:", d.List())

	
	
}
