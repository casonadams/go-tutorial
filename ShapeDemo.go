package main

import (
	"fmt"

	shape "github.com/casonadams/shapes"
)

func main() {
	r := shape.NewRectangle(9, 6)
	t := shape.NewTriangle(3, 6)

	fmt.Println("Area of myRectangle: ", shape.ShapeArea(r))
	fmt.Println("Area of myTriangle: ", shape.ShapeArea(t))

}
