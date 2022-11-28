package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter the num: ")
	fmt.Scanf("%d", &num)
	if n := num / 2; num%2 == 0 {
		fmt.Printf("%d is even, %d times present", num, n)
	} else {
		fmt.Printf("%d is odd, %d times present", num, n)

	}
}
