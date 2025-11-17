package main

import (
	"fmt"

	"github.com/headfirstgo/magazine"
)

func main() {
	var s1 magazine.Subscriber
	var e1 magazine.Employee
	e1.Name = "이인하"
	e1.Address.City = "인천"
	s1.Address.City = "서울"
	e1.Salary = 100000
	fmt.Println(e1.Name, e1.Salary)
	fmt.Println(s1.Name, s1.Address.City)
}
