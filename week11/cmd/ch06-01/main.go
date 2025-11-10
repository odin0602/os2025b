package main

import "fmt"

func main() {

	subjects := []string{"Go", "Javascript", "Python", "Linux"} //slice 리터럴
	subjectsSlice := subjects[1:3]                              //slicing

	for _, subject := range subjects {
		fmt.Println(subject)
	}
	fmt.Println("===========================================")
	for i := 0; i < len(subjectsSlice); i++ {
		fmt.Println(subjects)
	}
}
