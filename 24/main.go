package main

import (
	"fmt"
	"math"
)

/*
Разработать программу нахождения расстояния между двумя точками, которые представлены в виде
структуры Point с инкапсулированными параметрами x,y и конструктором.
*/
// Структуры точки
type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// метод нахождения
func (point *Point) DistanceBetweenPoints(a *Point) float64 {
	return math.Sqrt((a.x-point.x)*(a.x-point.x) + (a.y-point.y)*(a.y-point.y)) // по т. Пифагора
}
func main() {
	point := new(Point)
	point.x = 10
	point.y = 20
	point1 := new(Point)
	point1.x = 20
	point1.y = 10

	fmt.Println(point.DistanceBetweenPoints(point1))
}
