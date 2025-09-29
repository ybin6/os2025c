package main

import (
	"fmt"
	"reflect"
)

func main() {
	//var 64f float64  //error, must start with letter.
	var f64 float64
	// total_Price := 1000
	// totalprice := 1000
	totalPrice := 1000 //t는 소문자 될 수도 있고 대문자 될 수도 있다.

	fmt.Println(totalPrice)
	fmt.Println(f64, reflect.TypeOf(f64))

}
