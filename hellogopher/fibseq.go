package main

import (
	"fmt"
)

var previous int = 0
var current int = 0

func main() {
	fib(0, 100000)
}

func fib(num int, max int) {
	if num == 0 {
		fmt.Println(0)
		num = 1
	}
	if num <= max {
		current = num + previous
		fmt.Println(current)
		previous = num
		fib(current, max)
	}
}
