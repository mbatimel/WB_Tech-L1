package main

import "fmt"

func main(){
	data := []string{"cat", "cat", "dog", "cat", "tree"}
	result := make(map[string] bool)
	for _, v := range data{
		result[v] = true
	}
	// итерируемся по последовательности строк, устанавливаем строку в качестве
	// ключа карты result и присваиваем ключу значение в true, поскольку
	// карты не допускают дублирования ключей, получаем имитация множества.
	for v := range result{
		fmt.Printf("%s ", v)
	}
	fmt.Printf("\n")
}