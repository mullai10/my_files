package main

import "fmt"

func smallestEvenMultiple(n int) int {
    if n %2 == 0{
        return n
    }else {
        return 2*n
    }
    
}

func main(){

	fmt.Print(smallestEvenMultiple(5))
}