package greetingspackage

import "fmt"

func PrintGreetings() {
	fmt.Println("I'm printing a message from the PrintGreetings() function!")
}

func printGreetingsUnexported() {
	fmt.Println("I'm printing a message from the printGreetingsUnexported() function!")
}
