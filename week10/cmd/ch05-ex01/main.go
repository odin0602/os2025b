package main

import (
	"fmt"
)

func main() {
	//arrayBool := [3]bool{true, false, true} //array 리터럴
	numbers := [3]int{-9, 11, 7}
	for i, number := range numbers {
		fmt.Println(i, number)
	}
	//fmt.Printf("%#v\n", arrayBool)
	//fmt.Printf("%#v\n", arrayInt)
	//fmt.Println(reflect.TypeOf(arrayInt))
}
