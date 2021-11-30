package main

import (
	"fmt"
	"net/http"
	"strings"
)

func indexFunc (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "First page here!")
}

func startWeb (port string) {
	// remove any colon if present (just for print statement)
	port = strings.Trim(port, "s")

	fmt.Printf("Attempting to start webserver on port %v...", port)
	http.HandleFunc("/", indexFunc)

	// prepend the colon
	portArg := ":" + port
	http.ListenAndServe(portArg, nil)
}

func takeOptions () {
	fmt.Print("Start web server? (Y/N): ")

	var choice string
	fmt.Scanln(&choice)
	// make uppercase
	choice = strings.ToUpper(choice)

	switch choice {
	case "Y":
		fmt.Print("Please enter a port # for server to run: ")
		var port string
		fmt.Scanln(&port)
		portArg := strings.TrimSpace(port)
		startWeb(portArg)
	case "N":
		fmt.Println("You chose not to start.")
	default:
		fmt.Println("Invalid input")
		takeOptions()
	}
}

func main () {
	fmt.Println("Hello World")

	takeOptions()	
}