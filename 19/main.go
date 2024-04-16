package main

import (
	"fmt"
	
)

func main() {
	rs := []rune("главрыба")
	for v := range SortStart(rs){
		fmt.Printf("%s",string(rs[v]))
	}

}

func quickSort(arr []rune, low, high int) []rune {
	
	i := low
	for j := high; j > high/2; j-- {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		
	}
	
	return arr
}

func SortStart(arr []rune) []rune {
	
	return quickSort(arr, 0,len(arr)-1)
}
