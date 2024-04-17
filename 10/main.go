package main

import (
	"fmt"
	"math"
)

func getKey(value float64) int64{
	return int64(value - math.Mod(value, 10))
}
func main() {
	data := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	mp := make(map[int64] []float64)
	for _, temp := range data {
		
		key :=getKey(temp)
		mp[key] = append(mp[key], temp)
	}
	/*
			Задача решается одним циклом. Проходим по каждому элементу слайса,
			Создаем ключ из каждого температурного слайса на каждой итерации.
			Затем записываем в нашу хэш-таблицу, полученный из температурного значения, ключ
			и по нему добавляем значение в слайс.
			Далее значения просто записываются в слайс по существующему ключу, если он уже есть.
			При создании ключа температура делится на 10, округляется в меньшую сторону и вычитается из исходного значения
		*/
		
	
	for key, value := range mp {
		fmt.Println(key, value)
	}

}