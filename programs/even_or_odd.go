package main

import "fmt"

func main(){
	fmt.Print("Enter the Number: ")
	var num int
	fmt.Scanln(&num)
	
	if num %2 == 0{
		fmt.Println(num,"is even number")
	}else{
		fmt.Println(num,"is odd number")
	}	
	
}