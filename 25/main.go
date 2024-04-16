package main

import (
	"fmt"
	"time"
)

// В функцию sleep передаем продолжительность времени сна.
func Sleep(d time.Duration) {
	// Текущая горутина блокируется пока в канал созданный методом time.After
	// не передаться текущее время, после отсчета таймером продолжительности
	// времени, значение которой было передано в метод time.After.
	<-time.After(d)
}

func main() {
	// Получаем метку времени для отслеживания.
	beforeSleep := time.Now()
	// ВЫзываем функцию Sleep с 3 секундами.
	Sleep(3 * time.Second)
	// Печатаем сколько времени прошло с получения метки beforeSleep и
	// выполнения функции Sleep
	fmt.Println(time.Since(beforeSleep)) // ~ 3.00219791s
}