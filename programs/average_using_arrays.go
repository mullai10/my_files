package main

import "fmt"

func main() {

	fmt.Print("Enter the number of elements: ")
	var number int
	var user_input [100]int
	sum := 0 
	var average int
	
	fmt.Scanln(&number)
	
	
	for i:= 1; i<=number; i++{
		fmt.Print("Enter the number: ")
		fmt.Scanln(&user_input[i])
		sum = sum + user_input[i]
	}
	average = sum/number
	fmt.Println("average is",average)

}