package main

import "fmt"

func main() {
	//var arr = [5]int{2, 56, 34, 87, 568}
	arr2 := []int{76, 90, 34, 27, 34, 81, 986}
	// // for i := 0; i < len(arr2); i++ {
	// // 	fmt.Printf("%d\n", arr2[i])
	// }
	for _, n := range arr2 {
		fmt.Printf("%d\n", n)
	}

}
