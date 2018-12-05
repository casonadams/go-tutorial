package greetingspackage

import "fmt"

var MagicNumber int

func GopherGrettings() {
	fmt.Println("A very jolly hello my fellow gophers! I'm printing from the GopherGreetings() fucntion")
	printGreetingsUnexported()
}

func init() {
	MagicNumber = 108
}
