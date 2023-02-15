package main

import "fmt"

func main(){

	var num[100] int
	var temp int
	
	fmt.Print("Enter number of elements: ")
	fmt.Scanln(&temp)
	
	for i := 0; i < temp; i++ {
		fmt.Print("Enter the number : ")
		fmt.Scan(&num[i])
	}
	max := 0
	
	for j := 0; j < temp; j++ {
		if max < num[j]{
			max = num[j]
		}

	}
	

	fmt.Print("The largest number from the given array is : ",max)

}