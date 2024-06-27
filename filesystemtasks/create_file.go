package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("example.txt") // we are using os module to create the file and give path of the file (here example.txt)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	fmt.Println("File created successfully.")
}
