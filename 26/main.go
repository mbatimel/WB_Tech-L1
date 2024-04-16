package main

import (
	"fmt"
	"strings"

)

func isUnique(s string) bool {
	// Преобразуем строку в нижнему регистру.
	lowerStr := strings.ToLower(s)
	// Карта для подсчета символов в строке.
	rMap := map[rune]int{}
	// Итерируемся по строке, каждый символ используем в качестве ключа в карте
	// и инкерментируем значение этого ключа.
	for _, r := range lowerStr {
		rMap[r]++
	}
	// Итерируемся по карте, если значение ключа больше 1, значит строке есть
	// повторяющийся символ, в этом случае возвращаем false.
	for _, v := range rMap {
		if v > 1 {
			return false
		}
	}
	// Если в карте значения всех ключей меньше 2, то в строке все
	// символы уникальны, возвращаем true.
	return true
}

func main() {
	s1 := "Hi man😎"
	s2 := "🌚Hello world"
	s3 := "abcd"
	s4 := "abCdefAaf"
	s5 := "aabcd"

	fmt.Println(isUnique(s1))
	fmt.Println(isUnique(s2))
	fmt.Println(isUnique(s3))
	fmt.Println(isUnique(s4))
	fmt.Println(isUnique(s5))
}