package main

import "fmt"

type queue struct {
	array []int
}

func (q *queue) enqueue(data int) {
	q.array = append(q.array, data)
}

func (q *queue) size() int {
	return len(q.array)
}
func (q *queue) dequeue() {
	if len(q.array) == 0 {
		fmt.Println("empty!")
	}
	fmt.Println("the dequeue element is", q.array[0])
	q.array = q.array[1:]
}

func (q *queue) Display() {
	var res []int
	for i := range q.array {
		res = append(res, q.array[i])
	}
	fmt.Println(res)
}
func main() {

	//var que []int
	que := queue{}

	que.enqueue(10)
	que.enqueue(20)
	que.enqueue(30)
	que.enqueue(40)

	que.size()
	que.Display()

	que.dequeue()
	que.size()
	que.Display()
}
