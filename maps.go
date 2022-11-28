package main

import "fmt"

func main() {

	var a = make(map[string]string) // the map is empty now

	a["brand"] = "Ford"
	a["model"] = "mustang"
	a["year"] = "1964"

	fmt.Println(a)

	for key, value := range a {
		fmt.Printf("%v -> %v\n", key, value)
	}
}
