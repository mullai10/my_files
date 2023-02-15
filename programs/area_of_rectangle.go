package main

import "fmt"

func main(){

	var l, b float64
	
	var area float64
	
	fmt.Print("Enter the length of rectangle: ")
	fmt.Scan(&l)
	
	fmt.Print("Enter the breadth of rectangle: ")
	fmt.Scan(&b)
	
	area = l * b
	
	fmt.Println("area of rectangle is", area)
}
	