package main

import "fmt"

func main(){
	data := []int{1,2}
	data[0], data[1] = data[1], data[0]
	fmt.Print(data)
	one:=1
	tow:=2
	one, tow = tow,one
	fmt.Printf("\none: %d, two: %d\n",one,tow)


}