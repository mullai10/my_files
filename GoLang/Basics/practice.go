package main

import "fmt"

func main() {

	// 	//varaiables

	// 	/*var name string = "mullai nesan"
	// 	name = "mullai nesan"
	// 	var number int = 45
	// 	var num float32 = 28.10
	// 	name := "mullai nesan"
	// 	number := 23
	// 	num := 28.10

	// 	fmt.Printf("my name is %v and my age is %v\n", name, number)
	// 	fmt.Printf("number type is %T\n", num)
	// 	fmt.Println(num)*/

	// 	//arrays

	// 	/*grades := [3]int{}
	// 	grades := [...]int{45, 28, 10}
	// 	grades[2] = 17
	// 	var a = [5]int{12, 13, 14, 15, 16}
	// 	b := a
	// 	b[3] = 34

	// 	fmt.Printf("Grades: %v\n", grades[1])
	// 	fmt.Println("length of grades is", len(grades))
	// 	fmt.Printf("the values of a is %v\n", a)
	// 	fmt.Printf("the values of b is %v\n", b)*/

	// 	//slices

	// 	/*a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 	num1 := a[:]
	// 	num2 := a[2:]
	// 	num3 := a[:8]
	// 	num4 := a[3:9]
	// 	b := make([]int, 6)
	// 	c := []int{}
	// 	c = append(c, 3)
	// 	c = append(c, 2, 3, 4, 5)

	// 	fmt.Println(num1)
	// 	fmt.Println(num2)
	// 	fmt.Println(num3)
	// 	fmt.Println(num4)
	// 	fmt.Printf("length: %v\n", len(a))
	// 	fmt.Printf("capacity: %v\n", cap(a))
	// 	fmt.Println(b)
	// 	fmt.Println(c)*/

	// 	//pointers

	// 	/*var a int = 28
	// 	var ptr *int // using &symbol to store a address of the variable to pointer
	// 	ptr = &a
	// 	fmt.Println(ptr)  // to print the address of the pointor varaibale
	// 	fmt.Println(*ptr) //dereferences --> to print the value of the variable which is stored the address of that variable
	// 	fmt.Println(&a)*/

	// 	//maps

	// 	/*subject_marks := map[string]float32{"Tamil": 95, "English": 81, "Maths": 93, "Science": 100, "Social_Science": 95}
	// 	var subject_marks = map[string]float32{"Tamil": 95, "English": 81, "Maths": 93, "Science": 100, "Social_Science": 95}

	// 	subject_marks["English"] = 85
	// 	subject_marks["Total"] = 468

	// 	delete(subject_marks, "Total")

	// 	fmt.Println(subject_marks)
	// 	fmt.Println(subject_marks["Tamil"]) */

	// 	//struct

	// 	/*type person struct {
	// 		name string
	// 		age  int
	// 	}
	// 	person1 := person{"john", 25}
	// 	person2 := person{"paul", 27}

	// 	fmt.Println(person1)
	// 	fmt.Println(person2) */

	// 	/*type rectangle struct {
	// 		length int
	// 		breath int
	// 	}
	// 	rect := rectangle{28, 10}

	// 	fmt.Println("length: ", rect.length)
	// 	fmt.Println("breath: ", rect.breath)

	// 	area := rect.length * rect.breath

	// 	fmt.Println("area: ", area) */

	//maps

	var a = map[string]int{"mullai": 1, "dhinesh": 2, "praveen": 3}

	var b = make(map[string]int)
	for k, v := range a {
		fmt.Printf("%v : %v\n", k, v)
	}
	b["car"] = 4
	b["auto"] = 3
	b["bicycle"] = 2
	fmt.Println(b)
	b["lorry"] = 6

	delete(b, "bicycle")
	fmt.Println(b)

	var input string
	fmt.Scanf("%s", &input)
	wheels, ok := b[input]
	if !ok {
		fmt.Println("Key Not found")
		return
	}
	fmt.Printf("No of wheels: %d", wheels)

}
