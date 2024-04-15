package main

import "fmt"

func main() {
	var a int64 = 0
	var pos uint8 = 0
	var bit uint8 = 0
	fmt.Println("select the number you want to change")
	fmt.Scan(&a)	
	var old_a int64
	fmt.Println("select the bit position to change")
	fmt.Scan(&pos)
	fmt.Println("select 1 or 0")
	fmt.Scan(&bit)
	if pos > 64{
		fmt.Print("out of range")
		return
	}
	switch bit {
	case 0:
		var mask int64
		mask = ^(1<<pos)
		a &= mask
	case 1:
		a |= (1<<pos)	
	}
	fmt.Printf("old number %d, old bit mask %b\n", old_a, old_a)
	fmt.Printf("new number %d, new bit mask %b", a, a)
}

