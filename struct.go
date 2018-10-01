package main

import ("fmt")

type Rectangle struct {
	length, breadth int
}


func area(rec Rectangle) int {
	return rec.length * rec.breadth
}

func (r *Rectangle ) area2() int {
	r.length = 6
	return r.length * r.breadth
}

func main() {
	var rec Rectangle;
	fmt.Println(rec.length, rec.breadth)
	fmt.Println(area(rec))
	
	sec := new(Rectangle)
	sec.length = 5
	sec.breadth = 4
	fmt.Println(sec.length, sec.breadth)
	fmt.Println(area(*sec))

	fmt.Println(sec.area2())
	fmt.Println(area(*sec))
	fmt.Println(sec.length, sec.breadth)

}


