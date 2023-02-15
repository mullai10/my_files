package main

import "fmt"

func findNumbers(nums []int) int {
    count := 0
    for i:=0; i<len(nums); i++{
        j:= nums[i]
        if len(j) %2 ==0{
            count += 1
        }
    }
    return count
    
}

func main(){

	nums := []int {12,345,67,6,7689,456132, 78654}
	fmt.Print(findNumbers(nums))
}