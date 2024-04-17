package main

import (
	"fmt"
	"strings"
)

//сплитуем нашу строку по пробелам далее как делали в 19 номере
//просто меняем местами 
func main() {
	rs := "snow dog sun"
	fmt.Println(SortStart(rs))

}

func quickSort(arr []string, low, high int) []string {
	
	i := low
	for j := high; j > high/2; j-- {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		
	}
	return arr
}
func SortStart(str string) []string {
	split_str:= strings.Split(str," ")
	return quickSort(split_str, 0, len(split_str)-1)
}
