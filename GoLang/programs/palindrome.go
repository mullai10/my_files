package main

import "fmt"

func main() {

	var num, temp,remainder int
	
	fmt.Print("Enter the number:")
	fmt.Scan(&num)
	
	temp = num
	
	sum := 0
	
	for temp>0{
		remainder = temp % 10
		sum = (sum*10) + remainder
		temp = temp/10
	}
	
	if sum == num{
		fmt.Print("the given number is palindrome")
	}else{
		fmt.Print("the given number is not palindrome")
	}	
	
}