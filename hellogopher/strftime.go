package main

import (
	"fmt"
	"time"
)

func main() {
	timeString := "2018-12-20 13:44:22,456"
	timePattern := "%Y-%m-%d %H:%M:%S,%f"
	fmt.Println(Parse(timeString, timePattern))
}

func Parse(s string, p string) time.Time {

	return time.Now()
}
