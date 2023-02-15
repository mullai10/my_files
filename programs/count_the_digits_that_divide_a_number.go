package main

import "fmt"

func countDigits(num int) int {
    temp := num
    count := 0
    for temp > 0{
        rem := temp % 10
        if num%rem == 0 {
            count++
        }
        temp = temp/10
    }
    return count
    
}

func main(){
	fmt.Print(countDigits(32))
}
