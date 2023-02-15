package main

import "fmt"

func main() {
	var c byte
	fmt.Print("Do you want subscribe? say(y/n): ")
	fmt.Scanf("%c", &c)
	switch c {
	case 'y':
		fmt.Println("Thank You!")
	case 'n':
		fmt.Println("No Problem")
	default:
		fmt.Println("Invalid Input")

	}

}
