package main

import (
	"fmt"
	"strings"
)

func main () {
	fmt.Println("Hello World")
	takeOptions()
}

func takeOptions () {
	fmt.Print("Start web server? (Y/N): ")

	var choice string
	fmt.Scanln(&choice)
	// make uppercase
	choice = strings.ToUpper(choice)

	switch choice {
	case "Y":
		serverMenu()
		takeOptions() // if exited server menu, regive options
	case "N":
		fmt.Println("You chose not to start.")
	default:
		fmt.Println("Invalid input")
		takeOptions()
	}
}