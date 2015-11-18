package main

import "fmt"

type Rectangle struct {
	width, height int
}

func (r *Rectangle) getArea() int {
	return r.width * r.height
}

func (r *Rectangle) getCircumstance() int {
	return (r.width * 2) + (r.height * 2)
}

func main() {
	var firstRec = Rectangle{5, 5}
	fmt.Printf("Area of rec : %d\n", firstRec.getArea())
	fmt.Printf("Circumstance of rec : %d\n", firstRec.getCircumstance())
}