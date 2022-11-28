package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	str := "india"
	fmt.Println(str)
	fmt.Println(reflect.TypeOf(str))

	//length of the string
	fmt.Println(len(str))

	//ToUpper
	fmt.Println(strings.ToUpper(str))

	//ToLower
	fmt.Println(strings.ToLower(str))

	//HasPrefix
	fmt.Println(strings.HasPrefix(str, "in"))

	//HasSuffix
	fmt.Println(strings.HasSuffix(str, "ia"))

	//Repeat
	fmt.Println(strings.Repeat(str, 3))

}
