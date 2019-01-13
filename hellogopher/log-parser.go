package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
	file, err := os.Open("/tmp/foo")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// layout := "2006-01-02 15:04:05,000"
	timeArray := []int64{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		timeStamp := cleanTime(line)
		timeArray = append(timeArray, timeStamp)

		// fmt.Println(time.Parse(layout, scanner.Text()))
	}

	diff := (timeArray[1] - timeArray[0])
	fmt.Println(diff / 1000000)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func cleanTime(s string) int64 {
	re := regexp.MustCompile(`\d+-\d+-\d+ \d+:\d+:\d+,\d+`)
	s = re.FindString(s)
	s = strings.Replace(s, ",", ".", -1)
	s = strings.Replace(s, " ", "T", -1)
	s = s + "000000Z"
	t, _ := time.Parse(time.RFC3339Nano, s)
	nanoTime := t.UnixNano()
	return nanoTime
}
