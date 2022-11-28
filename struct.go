package main

import "fmt"

type Employee struct {
	name   string
	id     int
	salary int
}

func main() {

	//creating struct specifying field names

	emp1 := Employee{
		name:   "Mullai Nesan",
		id:     1001,
		salary: 25000,
	}

	//creating struct without specifying field names

	emp2 := Employee{"Abishek", 1002, 28000}

	fmt.Println("Employee 1:", emp1)
	fmt.Println("Employee 2:", emp2)

	//access the struct individual fields using . dot operator

	fmt.Println("Employee 1 id is", emp1.id)
	fmt.Printf("Employee 2 salary is %d", emp2.salary)

	//pointers to a struct

	emp3 := &Employee{"praveen", 1003, 30000}

	fmt.Println("\nEmployee 3:", (*emp3))
	fmt.Println("Employee 3 name is", &emp3.name)
	fmt.Println("Employee 3 name is", emp3.name)
	fmt.Println("Employee 3 name is", (*emp3).name)

}
