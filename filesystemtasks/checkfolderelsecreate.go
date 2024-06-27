package main

import (
	"fmt"
	"os"
)

func main() {
	path := "example_folder"

	// Check if file/folder exists usig OS module and stat function
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// Create folder if it doesn't exist
		err := os.Mkdir(path, 0755) //0755 - permission for owner to read write execute
		if err != nil {
			fmt.Println("Error creating folder:", err)
			return
		}
		fmt.Println("Folder created successfully.")
	} else {
		fmt.Println("Folder already exists.")
	}
}
