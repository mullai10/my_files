package main

import "fmt"

type node struct {
	data int
	next *node
}

type LinkedList struct {
	head   *node
	length int
}

//--------------Insertion at end----------------

func (l *LinkedList) Insert_end(val int) {
	//n := node{data : val}
	n := node{}
	n.data = val
	if l.head == nil {
		l.head = &n
		l.length++
		return
	}
	temp := l.head
	for i := 0; i < l.length; i++ {
		if temp.next == nil {
			temp.next = &n
			l.length++
			return
		}
		temp = temp.next
	}
}

//-----------insertion at front-----------

func (l *LinkedList) Insert_Front(val int) {
	n := node{data: val}
	if l.length == 0 {
		l.head = &n
		l.length++
		return
	}
	temp := l.head
	n.next = temp
	l.head = &n
	l.length++
	return
}

//-----------searching an value is present or not in that list------------

func (l *LinkedList) search(val int) {
	temp := l.head
	for i := 0; i < l.length; i++ {
		if temp.data == val {
			fmt.Printf("the value %d is there", temp.data)
			return
		}
		temp = temp.next
	}
	fmt.Println("the given value is not there!")
}

//----------Get the position of the element----------

func (l *LinkedList) Get_at(pos int) *node {
	temp := l.head
	for i := 0; i < pos; i++ {
		temp = temp.next
	}
	return temp

}

//-----------insertion at given position----------

func (l *LinkedList) Insert_At(pos int, val int) {
	n := node{data: val}
	if l.length == 0 {
		l.head = &n
		l.length++
		return
	}
	curr := l.Get_at(pos)
	n.next = curr
	prev := l.Get_at(pos - 1)
	prev.next = &n
	l.length++
}

//------------deletion at First-----------

func (l *LinkedList) Delete_First() {
	temp := l.head
	fmt.Println("\nDeleted Element:", temp.data)
	l.head = temp.next
	l.length--
	return
}

//----------deletion at end----------

func (l *LinkedList) Delete_End() {
	temp := l.head
	for i := 0; i < l.length-1; i++ {
		temp = temp.next
	}
	temp.next = nil
	l.length--
	return
}

//-------------deletion at given position--------

func (l *LinkedList) Delete_At(pos int) {
	if l.length == 0 {
		fmt.Println("No data found!")
		return
	}
	prev := l.Get_at(pos - 1)
	next := l.Get_at(pos + 1)
	prev.next = next
	l.length--
}

//-----------Display the given linked list-----------

func (l *LinkedList) Display() {
	if l.length == 0 {
		fmt.Println("No node is there")
	}
	temp := l.head
	for i := 0; i < l.length; i++ {
		fmt.Printf("%v ->", temp.data)
		temp = temp.next
	}
}

func main() {

	list := LinkedList{}

	list.Insert_end(10)
	list.Insert_end(20)
	list.Insert_end(30)

	list.Insert_Front(5)

	list.Display()

	fmt.Println()
	list.search(20)

	fmt.Println()
	list.Insert_At(2, 15)
	list.Display()
	fmt.Println()

	// list.Delete_First()
	// list.Display()

	// fmt.Println()
	// list.Delete_End()

	list.Delete_At(2)
	list.Display()
	// list.Display()
}
