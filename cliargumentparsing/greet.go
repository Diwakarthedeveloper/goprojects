package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 { // to check nothing useless is provided as argument
		fmt.Println("Please provide a name as an argument")
		//give the argument like go run greet.go "Diwakar Jha"
		return

	}
	name := os.Args[1] //

	fmt.Println("Hello", name, "how are you")
}
