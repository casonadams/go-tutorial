package greetingspackage

import "fmt"

func PrintGreetings() {
	fmt.Println("I'm printing a message from the PrintGreetings() function!")
}

func printGreatingsUnexported() {
	fmt.Println("I'm printing a message from the printGreetingsUnexported() function!")
}
