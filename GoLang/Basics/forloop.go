package main

import "fmt"

func main() {
	var temp, sum int
	sum = 0
	fmt.Print("Enter the numbers: \n")
	for i := 0; i < 5; i++ {
		fmt.Scanf("%d\n", &temp)
		sum = sum + temp
	}
	fmt.Printf("the sum is %d", sum)
}
