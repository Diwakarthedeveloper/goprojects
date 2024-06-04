package main

import (
	"fmt"
	"strings"
)

func main() {
	// concatination of strings
	a := "Hello, How are you Diwakar "
	b := "I am Diwakar"
	concat := a + " " + b
	fmt.Println("Greet = ", concat)

	//Splitting a String with , or space or any sumbol or word
	tod := strings.Split(concat, "_")
	fmt.Println("Sentence Broken with @ = ", tod)

	//Replace a substring
	badal := strings.Replace(concat, "Diwakar", "Mr DJ", 1) // 1 here denotes that to replace only first occurance of Diwakar
	fmt.Println("replacement= ", badal)

	// changing upper case to lower case and vice versa
	bada := strings.ToUpper(concat)
	chota := strings.ToLower(concat)
	fmt.Println("Uppercase = ", bada)
	fmt.Println("Lowercase= ", chota)

	// trimming strings
	trim := strings.TrimSpace(concat)
	fmt.Println("Trimmed string = ", trim)

	// substring - using sliceing method as go does not have default sub string method
	suru := 4
	aant := 14

	if aant <= len(concat) && suru < aant {
		kutout := concat[suru:aant]
		fmt.Println(" Substring: ", kutout)
	} else {
		fmt.Println("invalid range provided")
	}

	//checking if a string contains a substring
	checking := strings.Contains(concat, "Diwakar")
	fmt.Println("conatins substring = ", checking)

	//finding the substring of the baove indiex

	i := strings.Index(concat, "Diwakar")
	fmt.Println("Index of Diwakar", i)
}
