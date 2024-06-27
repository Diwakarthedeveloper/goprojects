package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("example.txt") // ioutil.ReadFile is used to read any file and its value in it and store in variable named "data"
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("File content:\n", string(data)) // the value in the file above is stored in data then anytype of data Type is converted here into string then only printed
}
