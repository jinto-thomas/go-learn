package main

import "fmt"

func main() {
	defer func() {
	str := recover()
		 fmt.Println(str)
		 fmt.Println("defer called..")
	}()
	fmt.Println(checkPanic(0))
}

func checkPanic(div int) float64 {
	return 10.0 / float64(div)
}
