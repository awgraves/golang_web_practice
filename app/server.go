package main

import (
	"fmt"
	"net/http"
	"strings"
)

func serverMenu () {
	fmt.Print("Please enter a port # for server to run: ")
	var port string
	fmt.Scanln(&port)
	portArg := strings.TrimSpace(port)

	startWeb(portArg)
}

func startWeb (port string) {
	// remove any colon if present (just for print statement)
	port = strings.Trim(port, "s")

	fmt.Printf("Attempting to start webserver on port %v...\n", port)
	http.HandleFunc("/", indexFunc)

	// prepend the colon
	portArg := ":" + port

	err := http.ListenAndServe(portArg, nil)

	if err != nil {
		fmt.Printf("\n%v\n\n", err)
		fmt.Println("Exiting program back to main menu...")
	}
}

func indexFunc (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "First page here!")
}