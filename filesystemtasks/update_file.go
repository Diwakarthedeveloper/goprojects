package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("example.txt", os.O_WRONLY|os.O_APPEND, 0644) // using openfile function from OS module and the mode in which we are opening Write_only or append
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // do the closing part at the end so defer used to get it done at th end

	if _, err := file.WriteString("\nNew content\n"); err != nil { // we are writing file (with New content)using file.WriteString if any error during this step , it will print below error
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("File updated successfully.") // if writting step is successfuull ot will print this
}
