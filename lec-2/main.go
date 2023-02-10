package main

import (
	"fmt"
)

type Dimension struct {
	length int
	width  int
	height int
}

func (d Dimension) Area() int {
	// fmt.Println(d.length * d.width * d.height)
	return d.length * d.width * d.height
}

func dimensions(length, width, height int) (area int) {

	return length * width

}

func main() {
	a, b := 1, 2
	n := &a
	fmt.Println(*n)
	*n = 3
	fmt.Println(a)
	o := &b
	fmt.Println(*o)
	*o = *o + 30
	fmt.Println(b)

	d := Dimension{5, 5, 5}
	x := dimensions(3, 3, 3)
	fmt.Println(d.Area())
	fmt.Println(x)

}
