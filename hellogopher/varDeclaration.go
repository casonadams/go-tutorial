package main

import "fmt"

func main() {
	// Declare a 32 bit integer without initializing it
	// Remember: varibles declared without an explicit initial value are given their zero value
	var myInteger int32

	// We expect to see a value of 0 since the zero-value of an int32 is 0
	fmt.Println("Value fo myInteger: ", myInteger)

	// Declare and Inialize 2 float64's and then show their sum
	var myFloatingPointNumber float64 = 36.333
	var myOtherFloatingPointNumber float64 = 306.969696969696
	fmt.Println("Sum of my floating point numbers: ", myFloatingPointNumber+myOtherFloatingPointNumber)

	// Go will figure out what the type is
	x, y, z := 0, 1, 2
	fmt.Printf("x: %d\ty: %d\tz: %d\n", x, y, z)

	myComplexNumber := 5 + 24i
	fmt.Println("Value of myComplexNUmber: ", myComplexNumber)

	//Example of grouping constant declarations/initializations
	const (
		speedOfLight     = 299792458 // unit: m/s
		pi               = 3.14
		myFavoriteNumber = 108
	)

	// Example of grouping variable declarations/initializations
	var (
		a = 0        // an integer
		b = 1.8 + 3i // a complex number
		c = 2.7      // a floating-point number
	)
	fmt.Printf("a: %v\tb: %v\tc: %v\n", a, b, c)
}
