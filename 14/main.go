package main

import (
	"fmt"
	"reflect"
	"strings"
)
func main() {
	var x interface{} = make(chan int)

	//первый способ определения типа переменных
	xType := fmt.Sprintf("%T", x)
	fmt.Println(xType)

	//второй способ	
	// Оператор switch использует специальное утверждение типа, в котором
	// используется ключевое слово type. Каждый оператор case указывает тип
	// и блок кода, который будет выполняться, когда значение, оцениваемое
	// оператором switch, имеет указанный тип. 
	switch v := x.(type) {
	case int:
		fmt.Println("int:", v)
	case float64:
		fmt.Println("float64:", v)
	case bool:
		fmt.Println("bool:", v)
	
	default:
		t := fmt.Sprintf("%T", v)
		if strings.Compare(t[0:4], "chan") == 0 {
			fmt.Println("Channel")
		} else {
			fmt.Println("...")
		}
	}


	//третий способ
	// Метод TypeOf из пакета reflect, возвращает тип переменной переданный в него.
	//	valueOf возвращает новое значение, инициализированное конкретным значением, хранящимся в интерфейсе i
	xType2 := reflect.TypeOf(x)
	xValue := reflect.ValueOf(x)
	fmt.Println(xType2, xValue)
}