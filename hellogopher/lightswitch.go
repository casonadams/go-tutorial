package main

import "fmt"

var lightSwitchIsOn = false

func main() {
	printLightSwitchState()
	toggleLightSwitch()
	printLightSwitchState()
	toggleLightSwitch()
	printLightSwitchState()
	toggleLightSwitch()
}

func printLightSwitchState() {
	fmt.Println("The light switch is off:", lightSwitchIsOn)
}

func toggleLightSwitch() {
	lightSwitchIsOn = !lightSwitchIsOn
}
