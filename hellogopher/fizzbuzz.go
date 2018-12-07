package main

import (
	"fmt"
)

var output string = ""

func main() {
	for i := 0; i < 50; i++ {
		fizzbuzz(i)
	}
}

func fizzbuzz(num int) {
	output = ""
	if num%3 != 0 && num%5 != 0 {
		fmt.Println(num, num)
		return
	}
	if num%3 == 0 {
		output += "fizz "
	}
	if num%5 == 0 {
		output += "buzz "
	}
	fmt.Println(num, output)
}
