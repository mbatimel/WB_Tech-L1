package  main

import "fmt"

func deleteFromSlice(s []int, i int) []int {
	// Присваиваем переменной slice значение входящего слайса
	// Используем append добавляя к слайсу до i-го элемента не включительно все значение после i-го элемента
	// имитируя тем самым удаление i-го элемента
	// "..." позволяет добавить сразу несколько элементов к слайсу
	if i >= len(s) {
		return s
	}
	return append(s[:i], s[i+1:]...)
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6}
	s = deleteFromSliceWithoutOrder(5,s)
	s2 := []int{0, 1, 2, 3, 4, 5, 6}
	s2 = deleteFromSlice(s2,5)
	fmt.Println(s)
	fmt.Println(s2)
}
// deleteFromSliceWithoutOrder также удаляет i-ый элемент из слайса, только с помощью перемещения последнего элемента
// на место i-го
func deleteFromSliceWithoutOrder(i int, s []int) []int {
	s[i] = s[len(s)-1]
	return append(s[:len(s)-1])
}