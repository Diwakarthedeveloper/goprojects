package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Remove("example.txt") // from OS module we will use Remove funtion to delete text from a particular path
	if err != nil {
		fmt.Println("Error deleting file:", err)
		return
	}

	fmt.Println("File deleted successfully.")
}
