package main

import "fmt"

func main() {

    var number,min,max,count int

    fmt.Print("Enter the Minimum Limit for Prime Numbers:")
    fmt.Scan(&min)
	
	fmt.Print("Enter the maximum Limit for Prime Numbers:")
	fmt.Scan(&max)

    fmt.Println("Prime Numbers between ", min, " and ", max, " are ")
    for number = min; number <= max; number++ {
        count = 0
        for i := 2; i < number/2; i++ {
            if number %i == 0 {
                count++
                break
            }
        }
        if count == 0 && number != 1 {
            fmt.Print(number)
        } 
    }
    fmt.Println()
}
