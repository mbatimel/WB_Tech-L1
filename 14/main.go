package main

import (
	"fmt"
	"reflect"
	"strings"
)
func main() {
	var x interface{} = make(chan int)
	xType := fmt.Sprintf("%T", x)
	fmt.Println(xType)

	
	switch v := x.(type) {
	case int:
		fmt.Println("int:", v)
	case float64:
		fmt.Println("float64:", v)
	case bool:
		fmt.Println("bool:", v)
	
	default:
		t := fmt.Sprintf("%T", v)
		if strings.Compare(t[0:4], "chan") == 0 {
			fmt.Println("Channel")
		} else {
			fmt.Println("...")
		}
	}


	
	xType2 := reflect.TypeOf(x)
	xValue := reflect.ValueOf(x)
	fmt.Println(xType2, xValue)
}