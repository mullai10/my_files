package users

import (
	"fmt"
	"strconv"
)

type User struct {
	id       int
	name     string
	password string
	age      int
	mail_id  string
	Balance  int
	next     *User
}

type Users_Details_List struct {
	head   *User
	length int
}

func (U *Users_Details_List) Add_User(name_val string, password_val string, age_val int, mail_id_val string) {

	new_user := User{name: name_val,
		password: password_val,
		age:      age_val,
		mail_id:  mail_id_val}

	if U.length == 0 {
		U.head = &new_user
		U.length++
		new_user.id = U.length
		return
	}
	temp := U.head
	//for temp != nil
	for i := 0; i < U.length; i++ {
		if temp.next == nil {
			temp.next = &new_user
			U.length++
			new_user.id = U.length
			return
		}
		temp = temp.next
	}
}

func (U *Users_Details_List) Display_Users() {
	if U.length == 0 {
		fmt.Println("There is no Users found!")
		return
	}
	temp := U.head
	fmt.Println("User Details:")
	for i := 0; i < U.length; i++ {
		fmt.Println("-------------------------")
		fmt.Printf("Name: %s,\nId: %d,\nPassword: %s,\nAge: %d,\nMail_id: %s,\n", temp.name, temp.id, temp.password, temp.age, temp.mail_id)
		fmt.Println("-------------------------")
		temp = temp.next
	}
	fmt.Println()
}

func (U *Users_Details_List) Check_User(name string, password string) string {
	if U.length == 0 {
		fmt.Println("No users found!")
		return "0 false"
	}
	temp := U.head
	for i := 0; i < U.length; i++ {
		if temp.name == name && temp.password == password {
			return strconv.Itoa(temp.id) + " true" + " " + temp.name
		}
		temp = temp.next
	}
	return "0 false"
}

func (u *Users_Details_List) Balance(id int) int {
	temp := u.head
	for i := 0; i < u.length; i++ {
		if id == temp.id {
			return temp.Balance
		}
		temp = temp.next
	}
	return 0
}

func (u *Users_Details_List) Recharge(id int, amount int) {
	if u.length == 0 {
		fmt.Println("There is no History found!")
		return
	}
	temp := u.head
	fmt.Println("History Details:")
	for i := 0; i < u.length; i++ {
		if id == temp.id {
			temp.Balance += amount

		}
		temp = temp.next
	}
}
