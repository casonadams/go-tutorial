package main

import (
	"fmt"
	"time"
)

func main() {
	loopTimer := time.NewTimer(time.Second * 9)
	for {
		fmt.Println("Inside the loop")
		<-loopTimer.C
		break
	}
}
