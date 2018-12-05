package main

import "fmt"

func main() {
	subject := "Gopher"

	fmt.Println("First element of Gopher string: ", string("Gopher"[0]))
	fmt.Printf("The first value of the subject string: %v\n", string(subject[0]))
	fmt.Printf("The last value of the subject string: %v\n", string(subject[len(subject)-1]))
	fmt.Println("Hello " + subject + "!")
}
