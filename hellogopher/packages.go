package main

import (
	"fmt"
	"github.com/casonadams/go-tutorial/hellogopher/volume1/section2/greetingspackage"
)

func main() {
	greetingspackage.PrintGreetings()
	greetingspackage.GopherGreetings()

	fmt.Println("The value of the Magic Number is:", greetingspackage.MagicNumber)
}
