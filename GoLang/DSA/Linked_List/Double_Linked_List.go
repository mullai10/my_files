package main

import "fmt"

type node struct {
	data int
	next *node
	prev *node
}

type Double_Linked_List struct {
	head   *node
	tail   *node
	length int
}

//---------Insertion at End--------

func (l *Double_Linked_List) Insert_End(val int) {
	n := node{data: val}
	if l.length == 0 {
		l.head = &n
		l.tail = &n
		l.length++
		return
	}
	temp := l.tail
	n.prev = temp
	temp.next = &n
	l.tail = &n
	l.length++

	// temp := l.head
	// for i := 0; i < l.length; i++ {
	// 	if temp.next == nil {
	// n.prev = temp
	// 		temp.next = &n
	// 		l.tail = &n
	// 		l.length++
	// 		return
	// 	}
	// 	temp = temp.next
	// }
}

//------------Insertion at Front------------

func (l *Double_Linked_List) Insert_Front(val int) {
	n := node{data: val}
	if l.length == 0 {
		l.head = &n
		l.tail = &n
		l.length++
		return
	}
	temp := l.head
	temp.prev = &n
	n.next = temp
	l.head = &n
	l.length++
	return
}

//---------Get an element of given position-------

func (l *Double_Linked_List) Get_At(pos int) *node {
	if pos < 0 {
		return l.head
	}
	temp := l.head
	for i := 0; i < pos; i++ {
		temp = temp.next
	}
	return temp
}

//---------Insertion at given position--------

func (l *Double_Linked_List) Insert_At(pos int, val int) {
	n := node{data: val}
	if l.length == 0 {
		l.head = &n
		l.tail = &n
		l.length++
		return
	}
	curr := l.Get_At(pos)
	prev := l.Get_At(pos - 1)
	n.next = curr
	n.prev = prev
	prev.next = &n
	curr.prev = &n
	l.length++
}

//-----------search an element---------

func (l *Double_Linked_List) search(val int) {
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

//-----------Deletion at front----------

func (l *Double_Linked_List) Delete_Front() {
	if l.length == 0 {
		fmt.Println("list is empty")
		return
	}
	temp := l.head
	fmt.Println("Deleted element is ", temp.data)
	l.head = temp.next
	temp.prev = nil
	l.length--
}

//---------Deletion at end-----------

func (l *Double_Linked_List) Delete_End() {
	if l.length == 0 {
		fmt.Println("Empty List!")
		return
	}
	temp := l.tail
	l.tail = temp.prev
	l.tail.next = nil
	l.length--
	return
}

//------------Deletion at given position----------

func (l *Double_Linked_List) Delete_At(pos int) {
	if l.length == 0 {
		fmt.Println("Empty list!")
		return
	}
	prev := l.Get_At(pos - 1)
	next := l.Get_At(pos + 1)
	prev.next = next
	next.prev = prev
	l.length--
}

//---------Display the elements of Linked list-----------

func (l *Double_Linked_List) Display() {
	if l.length == 0 {
		fmt.Println("Empty List!")
	}
	temp := l.head
	for i := 0; i < l.length; i++ {
		fmt.Printf("%d -> ", temp.data)
		temp = temp.next
	}
}

//---------Reverse the elements in Linked List--------

func (l *Double_Linked_List) Reverse() {
	if l.length == 0 {
		fmt.Println("Empty List!")
	}
	temp := l.tail
	for i := 0; i < l.length; i++ {
		fmt.Printf("<- %d ", temp.data)
		temp = temp.prev
	}

}
func main() {

	list1 := Double_Linked_List{}

	list1.Insert_End(10)
	list1.Insert_End(20)
	list1.Insert_End(30)

	list1.Display()
	fmt.Println()

	list1.Insert_At(1, 15)
	list1.Display()
	fmt.Println()

	list1.Reverse()
	fmt.Println()

	//list1.Delete_Front()
	//list1.Display()

	//list1.Delete_End()
	//list1.Display()

	list1.search(10)
	fmt.Println()

	list1.Delete_At(1)
	list1.Display()
}
