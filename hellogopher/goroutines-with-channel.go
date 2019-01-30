package main

import (
	"fmt"
	"math/rand"
	"time"
)

var done chan bool = make(chan bool)

func printGreetings(source string) {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	for i := 0; i < 9; i++ {
		w := r.Intn(1000)
		fmt.Println("Hello Gopher!", i, source, w)
		time.Sleep(time.Duration(w) * time.Millisecond)
	}

	if source == "goroutine" {
		done <- false // This can be true or false just needs to be sent
	}
}

func main() {
	go printGreetings("goroutine")
	printGreetings("main function")

	<-done
}
