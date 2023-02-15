package main

import "fmt"

type stack struct {
	array []int
}

func (s *stack) push(data int) {
	s.array = append(s.array, data)
}

func (s *stack) size() int {
	return len(s.array)
}

func (s *stack) pop() {
	if s.size() == 0 {
		fmt.Println("empty!")
	}
	fmt.Println("the pop element is", s.array[s.size()-1])
	s.array = s.array[0 : s.size()-1]
	fmt.Println()
}

func (s *stack) Display() {
	for i := 0; i < len(s.array); i++ {
		fmt.Println(s.array[i])
	}
	fmt.Println()
}

func main() {
	obj := stack{}

	obj.push(10)
	obj.push(20)
	obj.push(30)
	obj.push(40)

	obj.Display()

	obj.pop()

	obj.Display()

}
