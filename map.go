package main

import "fmt"

func main() {

	x := make(map[string]int)
	x["one"] = 1
	x["two"] = 2
	fmt.Println(x)
	delete(x, "one")
	name, ok := x["one"]
	fmt.Println(ok, name)
	fmt.Println(x)
}
