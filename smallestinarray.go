package main

import "fmt"

func main() {
	arr := []int {
			 45,
			 65,
			 63,
			 12,
			 52,
			 66,
			 7,
		 }

	var small int = arr[0]
	for _, val := range arr {
		if small > val {
			small = val
		}
	}

	fmt.Println(small)
}
