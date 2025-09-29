package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf(2.31))
	fmt.Println(reflect.TypeOf("go developer~"))
	fmt.Println(reflect.TypeOf('a'))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf(91))
	var gpa float32 = 3.99
	var name string = "kim inha"
	var id int = 1000
	var f64 float64
	var i16 int16
	var t bool
	var s string
	var i int

	fmt.Println(name, reflect.TypeOf(name))
	fmt.Println(id, reflect.TypeOf(id))
	fmt.Println(gpa, reflect.TypeOf(gpa))
	fmt.Println(f64, reflect.TypeOf(f64))
	fmt.Println(i16, reflect.TypeOf(i16))
	fmt.Println(t, reflect.TypeOf(t))
	fmt.Println(s, reflect.TypeOf(s))
	fmt.Println(i, reflect.TypeOf(i))
}
