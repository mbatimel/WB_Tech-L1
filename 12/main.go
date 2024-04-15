package main

import "fmt"

func main(){
	data := []string{"cat", "cat", "dog", "cat", "tree"}
	result := make(map[string] bool)
	for _, v := range data{
		result[v] = true
	}
	for v := range result{
		fmt.Printf("%s ", v)
	}
	fmt.Printf("\n")
}