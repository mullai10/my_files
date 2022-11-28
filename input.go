package main

import (
	"fmt"
)

func main() {

	//integer input from user
	//var i, j int

	//fmt.Print("Type two numbers: ")
	//fmt.Scan(&i, &j)
	//fmt.Scanf("%v %v", &i, &j)
	//fmt.Scanln(&i, &j)
	//fmt.Println("Your numbers are:", i, "and", j)

	//float input from user
	//var n1, n2 float64

	//fmt.Print("Type two Float numbers: ")
	//fmt.Scanf("%g %g", &n1, &n2)

	//fmt.Printf("Your Float numbers are: %g and %g", n1, n2)

	//string input from user

	//var fname, lname string

	//fmt.Print("Enter the fname and lname: ")
	//fmt.Scan(&fname, &lname)
	//fmt.Scanf("%s %s", &fname, &lname)
	//fmt.Scanln(&fname, &lname)

	//fmt.Printf("Your Full Name is %s %s", fname, lname)

	//bool input from user

	//var sunny bool

	//fmt.Println("Is the day sunny?")
	//fmt.Scanf("%t", &sunny)

	//fmt.Printf("%t, The day is sunny", sunny)

	//array input from user using loop

	fmt.Print("Enter the length of an array: ")
	var size int
	fmt.Scanln(&size)

	var names = make([]string, size)

	for i := 0; i < size; i++ {
		fmt.Printf("Enter the %dth element:", i)
		fmt.Scanf("%s", &names[i])
	}
	fmt.Println("Your array is:", names)
}
