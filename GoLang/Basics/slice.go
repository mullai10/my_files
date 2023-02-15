package main

import "fmt"

func main() {
	daysofweek := [...]string{
		"sun", "mon", "tue", "wed", "thu", "fri", "sat",
	}
	weekdays := daysofweek[1:6]
	weekdays[3] = "thursday"
	printslice(daysofweek[:])
	printslice(weekdays)
	numbers := make([]int, 5, 10)
	numbers = append(numbers, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Printf("%v\n", numbers)
	fmt.Printf("%d\n", len(numbers))
	fmt.Printf("%d\n", cap(numbers))

}

func printslice(slice []string) {
	fmt.Printf("%v\n", slice)
}
