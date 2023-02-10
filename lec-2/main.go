package main

import (
	"fmt"
)

type Dimension struct {
	length int
	width  int
	height int
}

func (d *Dimension) Area() int {
	//when we pass pointer the orginal values also effecetd

	// fmt.Println(d.length * d.width * d.height)
	d.height = 6

	return d.length * d.width * d.height
}

func dimensions(length, width, height int) (area int) {

	return length * width

}

func main() {
	d := Dimension{5, 5, 5}
	fmt.Println(d)
	fmt.Println(d.Area())
	fmt.Println(d)

}
