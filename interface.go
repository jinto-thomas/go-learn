package main

import ("fmt";"math")

type Calculator interface {
	area() float64
}

type Square struct {
	length float64
}

type Circle struct {
	radius float64
}

func (square *Square) area() float64 {
	return square.length * square.length
}

func (circle *Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

func printArea(cal Calculator) {
	fmt.Println(cal.area())
}

func main() {
	sq := Square{3}
	ci := Circle{3}

	printArea(&sq)
	printArea(&ci)
}
