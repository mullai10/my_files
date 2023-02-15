package main

import "fmt"

type node struct {
	data int
	next *node
}
type stack_list struct {
	length int
	head   *node
}

func (s *stack_list) Len() int {
	return s.length
}

func (s *stack_list) push(val int) {
	n := node{data: val}
	if s.length == 0 {
		s.head = &n
		s.length++
		return
	}
	temp := s.head
	for i := 0; i < s.length; i++ {
		if temp.next == nil {
			temp.next = &n
			s.length++
			return
		}
		temp = temp.next
	}
}

func (s *stack_list) Display() {
	if s.length == 0 {
		fmt.Println("Empty Stack!")
	}
	temp := s.head
	for i := 0; i < s.length; i++ {
		fmt.Printf("%d\n", temp.data)
		temp = temp.next
	}
	fmt.Println()
}

func (s *stack_list) pop() {
	if s.length == 0 {
		fmt.Println("Empty Stack!")
		return
	}
	temp := s.head
	for i := 0; i < s.length-1; i++ {
		temp = temp.next
	}
	temp.next = nil
	s.length--
	return
}
func main() {

	sta := stack_list{}

	sta.push(10)
	sta.push(20)
	sta.push(30)
	sta.push(40)

	sta.Display()

	sta.pop()
	sta.Display()

	fmt.Println(sta.Len())
}
