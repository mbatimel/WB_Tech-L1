package main

import "fmt"

type Human struct {
	FullName string
	
}
type Action struct {
	Human // Встроели тип Human, 
	//как анонимное поле(то есть переменная типа Human назвкана как Human) 
	//Так же мы могли могли "наследовать" структуру Human, как: Human Human )
	ActionType string
}

func main()  {
	a:= Action{Human{"Mangushev Ildar Nadirovich"},"running"}
	fmt.Printf("%s is %s",a.Human.FullName,a.ActionType)
}