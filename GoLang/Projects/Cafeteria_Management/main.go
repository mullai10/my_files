package main

import (
	"Cafeteria_Management/users"
	"fmt"
	"strconv"
	"strings"
)

func main() {

	fmt.Println()
	fmt.Println("* * * * * * * * * * * * * * * * * * * * * ")
	fmt.Println(" ---> Welcome to Our Cafeteria World! <--- ")
	fmt.Println("* * * * * * * * * * * * * * * * * * * * * ")

	var options int
	var Exit string

	user_list := users.Users_Details_List{}
	history_list := users.History_List{}

	fmt.Println()
	fmt.Println("Choose One Options!")
	fmt.Println()

	for ok := true; ok; ok = Exit == "y" {

		fmt.Println("1.Login 2.SignUp 3.Exit")

		fmt.Scanln(&options)

		switch options {
		case 1:
			fmt.Println()
			fmt.Println("You are now login into our Cafeteria world!")
			fmt.Println()
			Login(&user_list, &history_list)

		case 2:
			Sign_Up(&user_list)

		case 3:
			fmt.Println("Thank You! Please come again..")

		default:
			fmt.Println("Please pick correct options...!")
		}

		fmt.Print("Do you want to continue (y/n): ")
		fmt.Scanln(&Exit)
	}
}

func Sign_Up(users_list *users.Users_Details_List) {

	var name, mail_id, password string
	var age int

	fmt.Print("Enter your Name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter your Password: ")
	fmt.Scanln(&password)

	fmt.Print("Enter your Age: ")
	fmt.Scanln(&age)

	fmt.Print("Enter your Mail_id: ")
	fmt.Scanln(&mail_id)

	fmt.Println()

	users_list.Add_User(name, password, age, mail_id)

	users_list.Display_Users()

}

func Login(users_list *users.Users_Details_List, History *users.History_List) {

	var c_name, c_password string

	fmt.Print("Enter your Name: ")
	fmt.Scanln(&c_name)

	fmt.Print("Enter your Password: ")
	fmt.Scanln(&c_password)

	fmt.Println()

	success_msg := strings.Split(users_list.Check_User(c_name, c_password), " ")
	id, _ := strconv.Atoi(success_msg[0])
	user_name := success_msg[2]

	if success_msg[1] == "true" {
		fmt.Println("Successfully Logged in!")
		fmt.Println()
		After_Login(users_list, id, History, user_name)
	} else {
		fmt.Println("Please Enter Correct Details")
	}
	fmt.Println()
}

func After_Login(users_list *users.Users_Details_List, id int, History *users.History_List, user_name string) {

	fmt.Println("Choose your need:")
	fmt.Println()

	var need int
	var Back string

	for ok := true; ok; ok = Back == "y" {

		fmt.Println("1.Menu Card 2.Balance 3.Recharge 4.Your Orders 5.History")

		fmt.Scanln(&need)

		switch need {
		case 1:
			fmt.Println("You are entered now with menu card details")
			Menu_Card_List(History, id, user_name)

		case 2:
			balance := users_list.Balance(id)
			fmt.Printf("Your balance is %d\n", balance)

		case 3:
			reAmount := 0
			fmt.Println("How much you want to reacharge : ")
			fmt.Scanln(&reAmount)
			users_list.Recharge(id, reAmount)
			fmt.Println("Your recharge successfully processed..!")

		case 4:
			fmt.Println("your current orders is...!")
			History.Current_Order_History(id, users_list)

		case 5:

			History.Display_Order_History(id)

		default:
			fmt.Println("Please pick correct need!")

		}
		fmt.Println("Do you want go back (y/n):")
		fmt.Scanln(&Back)

	}
}

func Menu_Card_List(History_List *users.History_List, id int, username string) {

	var menus int
	var quantity int

	fmt.Println()
	fmt.Printf("1.Tea --> Rs.10\n2.Coffee --> Rs.15\n3.Green Tea --> Rs.15\n4.Lemon Tea -->Rs.15\n")
	fmt.Println()

	fmt.Print("Choose What you want? --> ")
	fmt.Scanln(&menus)

	fmt.Print("No Of Quantity --> ")
	fmt.Scanln(&quantity)

	switch menus {
	case 1:
		bill := Bill_Calculation(10, quantity)
		History_List.Add_History(id, username, "Tea", bill)

	case 2:
		bill := Bill_Calculation(15, quantity)
		History_List.Add_History(id, username, "Coffee", bill)

	case 3:
		bill := Bill_Calculation(15, quantity)
		History_List.Add_History(id, username, "Green Tea", bill)

	case 4:
		bill := Bill_Calculation(15, quantity)
		History_List.Add_History(id, username, "Lemon Tea", bill)

	default:
		fmt.Println("Choose Correct One!")
	}
}

func Bill_Calculation(cost int, quantity int) int {
	return cost * quantity

}
