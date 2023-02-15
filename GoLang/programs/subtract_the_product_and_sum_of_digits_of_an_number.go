package main

import "fmt"

func subtractProductAndSum(n int) int {
    sum := 0
    pro := 1
    for n>0 {
        rem := n % 10
        sum += rem
        pro *= rem
        n = n/10
    }
    return pro - sum
}

func main(){
	fmt.Print(subtractProductAndSum(234))
}