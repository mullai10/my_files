package main

import "fmt"

func getConcatenation(nums []int) []int {

    res := []int{}
	
    for i:=0;i<2;i++{
        for j:=0; j<len(nums); j++{
            res = append(res, nums[j])
        }
    }
    return res    
}

func main(){

	nums := [] int{1,2,1}
	fmt.Print(getConcatenation(nums))

}