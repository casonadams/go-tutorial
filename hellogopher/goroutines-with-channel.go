package main

import (
	"fmt"
	"math/rand"
	"time"
)

var done chan bool = make(chan bool)

func printGreetings(source string) {
	for i := 0; i < 9; i++ {
		r := rand.Intn(10000)
		fmt.Println("Hello Gopher!", i, source, r)
		time.Sleep(time.Duration(r) * time.Millisecond)
	}

	if source == "goroutine" {
		done <- true
	}
}

func main() {
	go printGreetings("goroutine")
	printGreetings("main function")

	<-done
}
