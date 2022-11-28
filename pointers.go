package main

import "fmt"

func main() {
	var a int
	var b *int
	a = 10                                         // both a and b are same
	b = &a                                         // a represent value of a and b represent address of a
	fmt.Printf("values of a, b is %v,%v\n", a, b)  //shows the address
	fmt.Printf("values of a, b is %v,%v\n", a, *b) // shows the value
	// aasign new value to pointer
	*b = 5
	fmt.Printf("values of a, b is %v,%v\n", a, b)
	fmt.Printf("values of a, b is %v,%v\n", a, *b)
}
