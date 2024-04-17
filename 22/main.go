package main

import (
	"fmt"
	"math/big"
)
// Для операций сложения, вычитания, умножения и деления чисел больше 2^20,
// используем пакет math/big, который позволяет оперировать большими числами.
// 1. big.Int для крупных целых чисел, когда 18 квинтиллионов недостаточно;
// 2. big.Float для вещественных чисел с плавающей запятой производной точности;
// 3. big.Rat для дробей вроде 1/3.

func main() {
	a := new(big.Int)
	b := new(big.Int)

	aStr := ""
	bStr := ""
	for {
		fmt.Printf("Enter first number: ")
		fmt.Scan(&aStr)
		// Создаем числа с помощью пакета big
		if _, ok := a.SetString(aStr, 10); ok {
			break
		}
		fmt.Println("Wrong data")
	}
	for {
		fmt.Printf("Enter second number: ")
		fmt.Scan(&bStr)
		// Создаем числа с помощью пакета big
		if _, ok := b.SetString(bStr, 10); ok {
			break
		}
		fmt.Println("Wrong data!")
	}
	fmt.Println(arifmeticOperation(a,b))

}
func arifmeticOperation(a, b *big.Int) *big.Int{
	result := new(big.Int)
	status := false
	for !status {
		operation := ""
		fmt.Printf("Select an operation (+, -, *, /): ")
		fmt.Scan(&operation)
		status = true
		switch operation {
		case "+":
			result.Add(a, b)
		case "-":
			result.Sub(a, b)
		case "*":
			result.Mul(a, b)
		case "/":
			zero := big.NewInt(0)
			if b.Cmp(zero) == 0 {
				fmt.Println("Divison by zero, select another operation")
				status = false
			} else {
				result.Div(a, b)
			}
		default:
			fmt.Println("Wrong operation!")
			status = false
		}
	}
	return result
}