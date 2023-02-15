package main

import "fmt"

func main(){
	var number, temp, remainder, result int
	fmt.Print("Enter the three digit number: ")
	fmt.Scan(&number)
	
	temp = number
	if number>99 && number<1000{
		for {
			remainder = temp % 10
			result += remainder * remainder * remainder
			temp = temp / 10
		
			if (temp == 0){
				break
			}
		}	
		
		if result == number{
			fmt.Print("the given number is armstrong number")
		}else{
			fmt.Print("the given number is not armstrong number")
			
		}
	}else{
		fmt.Print("please give 3 digit number")
	}
	
	
}
		
		
		