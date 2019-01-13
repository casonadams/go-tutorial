package main

import "fmt"

func nationCapitalsExample() {
	var nationCapitals map[string]string = make(map[string]string)
	nationCapitals["Afganistan"] = "Kabul"

	fmt.Println(nationCapitals)
}

func main() {
	nationCapitalsExample()
}
