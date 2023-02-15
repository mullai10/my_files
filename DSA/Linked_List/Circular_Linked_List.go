package main

import "fmt"

type c_node struct {
	data int
	next *c_node
}

type circular_Linked_List struct {
	head   *c_node
	length int
}

func (l *circular_Linked_List) Insert_end(val int) {
	c := c_node{data: val}
	if l.length == 0 {
		l.head = &c
		l.length++
		return
	}
	temp := l.head
	for i := 0; i < l.length; i++ {
		if temp.next == nil {
			temp.next = &c
			l.length++
		}
		temp = temp.next
	}
}

func (l *circular_Linked_List) Display() {
	if l.length == 0 {
		fmt.Println("Empty List!")
		return
	}
	temp := l.head
	for i := 0; i < l.length; i++ {
		fmt.Printf("%d ->", temp.data)
		temp = temp.next
	}

}
func main() {

	list := circular_Linked_List{}

	list.Insert_end(10)
	list.Insert_end(20)
	list.Insert_end(30)
	list.Display()

}
