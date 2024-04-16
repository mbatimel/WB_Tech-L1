package main

import "fmt"

/*var justString string
Потенциальный утечка памяти: Создается большая строка в функции createHugeString,
но после этого мы используете только первые 100 символов и присваиваете их переменной justString. 
Однако, оригинальная большая строка не освобождается, что может привести к утечке памяти.

Изменение глобальной переменной: Изменение глобальной переменной justString в 
функции someFunc без явного указания на это может усложнить отслеживание состояния программы и сделать код менее читаемым.
*/

func createHugeString(size uint64) string {
  fmt.Printf("Create new clise with size %d\n", size)
	return string(make([]rune, size))
}

func getSlice(v string, from, to uint64) string {
	data := []rune(v)
	slice := make([]rune, 0)
	for i := from; i < to; i++ {
		slice = append(slice, data[i])
	}
	return string(slice)
}

func someFunc() {
  var justString string
	v := createHugeString(1 << 10)
	justString = getSlice(v, 0, 100)

  fmt.Printf("Clise with size %d\n",len(justString))
}

func main() {
	someFunc()
}
