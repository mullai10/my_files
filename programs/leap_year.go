package main

import "fmt"

func main() {

    var year int

    fmt.Print("Enter the Year to Check for Leap = ")
    fmt.Scan(&year)

    if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
        fmt.Println(year, " is a Leap Year")
    } else {
        fmt.Println(year, " is Not a Leap Year")
    }
}