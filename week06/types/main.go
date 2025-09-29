package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Zero values - 할당하지 않았을 때 기본적으로 가지고 있는 값
	var f64 float64
	var t bool
	var s string
	var i int
	var i16 int16

	fmt.Println(f64, reflect.TypeOf(f64)) //수치형은 0
	fmt.Println(t, reflect.TypeOf(t))     //false
	fmt.Println(s, reflect.TypeOf(s))     //빈문자열
	fmt.Println(i, reflect.TypeOf(i))
	fmt.Println(i16, reflect.TypeOf(i16))
}
