package main

import "fmt"

type Rectangle struct {
	Width, Height int
}

func (r Rectangle) Area() int { // Value receiver
	return r.Width * r.Height
}

func (r *Rectangle) Scale(factor int) { // Pointer receiver
	r.Width *= factor
	r.Height *= factor
}

func recibir() {
	rect := Rectangle{Width: 10, Height: 5}
	area := rect.Area() // Calling the Area method
	fmt.Println("Area:", area)
	rect.Scale(2) // Calling the Scale method (modifies the original rect)
	fmt.Println("Scaled rect:", rect)
}
