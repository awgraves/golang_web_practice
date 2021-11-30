package main

import (
	"fmt"
	"strings"
)

func main () {
	fmt.Println("Hello World")
	mainOptions()
}

func mainOptions () {
	fmt.Println("S - to start web server")
	fmt.Println("F - to fetch web responses concurrently")
	fmt.Print(": ")

	var choice string
	fmt.Scanln(&choice)
	// make uppercase
	choice = strings.ToUpper(choice)

	switch choice {
	case "S":
		serverMenu()
		mainOptions() // if exited server menu, regive options
	case "F":
		fmt.Println("Would redirect to concurrent fetch")
	default:
		fmt.Println("Invalid input")
		mainOptions()
	}
}