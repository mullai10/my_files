package users

import "fmt"

type History struct {
	id         int
	name       string
	order_food string
	bill       int
	next       *History
}

type History_List struct {
	head *History
	len  int
}

func (H *History_List) Add_History(id_val int, name_val string, order_food_val string, bill_val int) {

	new_history := History{id: id_val, name: name_val, order_food: order_food_val, bill: bill_val}

	if H.len == 0 {
		H.head = &new_history
		H.len++
		return
	}
	temp := H.head
	for i := 0; i < H.len; i++ {
		if temp.next == nil {
			temp.next = &new_history
			H.len++
			return
		}
		temp = temp.next
	}
}

func (H *History_List) Display_Order_History(id int) {
	if H.len == 0 {
		fmt.Println("There is no History found!")
		return
	}
	temp := H.head
	fmt.Println("History Details:")
	for i := 0; i < H.len; i++ {
		if id == temp.id {
			fmt.Println("---------------------")
			fmt.Printf("Id: %d\nName: %s\nOrder_Food: %s\nBill: %d\n", temp.id, temp.name, temp.order_food, temp.bill)
			fmt.Println()
		}
		temp = temp.next
	}
}

func (H *History_List) Current_Order_History(id int, user_list *Users_Details_List) {

	var confirmation string
	if H.len == 0 {
		fmt.Println("There is no History found!")
		return
	}
	temp := H.head
	curr := H.head
	fmt.Println("History Details:")
	for i := 0; i < H.len; i++ {
		if id == temp.id {
			fmt.Println("---------------------")
			curr = temp
		}
		temp = temp.next
	}

	fmt.Printf("Id: %d\nName: %s\nOrder_Food: %s\nBill: %d\n", curr.id, curr.name, curr.order_food, curr.bill)
	fmt.Println()
	fmt.Println("Shall we confirm your order(y/n): ")
	fmt.Scanln(&confirmation)
	if confirmation == "y" {
		balance := user_list.Balance(id)
		if balance > curr.bill {
			user_list.Recharge(id, -curr.bill)
			fmt.Println("Your order is Confirmed..!")
		} else {
			fmt.Println("Please recharge first..!")
		}
	} else {
		fmt.Println("sorry to say,your order is cancelled..!")
	}

}
