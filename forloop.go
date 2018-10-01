package main

import "fmt"

func main() {
	fmt.Println("test run")
	
	var x [5] int
	for i:=1 ; i <= 5; i++{
		x[i-1] = i
	}

	y := [10]float64{34.1,12,43.4,12,5,1,53}
	fmt.Println(x)
	fmt.Println(y)

	var total int = 0
	for _, val := range x {
		total += val
	}

	fmt.Println("Total of x {%d}" , total)

	slice := y[0:5]
	fmt.Println(slice)

	a := make([]float64, 5)
	fmt.Println(a)
}
