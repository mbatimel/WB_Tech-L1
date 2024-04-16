package main

import (
	"fmt"
	"math"
)

// Структура для точки с полями x и y для обозначения координат точек.
type Point struct {
	x, y float64
}

// Функция конструктор точки.
func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

// Функция для рассчет расстояния между двумя точками.
// Расстояние между двумя точками равна корню квадратному из
// суммы квадратов разницы координат по каждой оси.
func DistanceSearch(p1, p2 *Point) float64 {
	// Вычисляем квадрат разницы координат по оси x.
	sqrtDivX := math.Pow(p1.x-p2.x, 2)
	// Вычисляем квадрат разницы координат по оси y.
	sqrtDivY := math.Pow(p1.y-p2.y, 2)
	// Вычисляем корень квадратный из суммы квадратов разницы координат
	// по каждой оси.
	return math.Sqrt(sqrtDivX + sqrtDivY)
}

func main() {
	firstPoint := NewPoint(2, 6)
	secondPoint := NewPoint(7, 9)
	fmt.Println(DistanceSearch(firstPoint, secondPoint)) // 5.830951894845301
}