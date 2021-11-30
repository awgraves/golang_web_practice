package main

import (
	"fmt"
	"strings"
)

func main () {
	fmt.Println("Hello World")
	mainOptions()
}

func getUserInput () string {
	// take from console
	var choice string
	fmt.Scanln(&choice)
	// make uppercase
	choice = strings.ToUpper(choice)

	return choice
}

func mainOptions () {
	fmt.Println("S - to start web server")
	fmt.Println("F - to fetch web responses concurrently")
	fmt.Println("Q - to quit this program")
	fmt.Print(": ")

	choice := getUserInput()

	switch choice {
	case "S":
		serverMenu()
		mainOptions() // if exited server menu, regive options
	case "F":
		fetchMenu()
		mainOptions()
	case "Q":
		fmt.Println("Quit")
	default:
		fmt.Println("Invalid input")
		mainOptions()
	}
}