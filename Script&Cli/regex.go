//regex is a very common operation used to validate the data. we use "regexp" model

package main

import (
	"fmt"
	"regexp"
)

func main() {

	name := "Diwakar Jha"
	email := "djha4635@gmail.com"
	phone := "+1-234-567-8900"

	// reger condyion pattern code
	namePattern := `^[a-zA-Z]+(?: [a-zA-Z]+)?$`
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	phonePattern := `^\+\d{1,3}-\d{3}-\d{3}-\d{4}$`

	// comparing regex condition with the input provided

	nameMatching, _ := regexp.MatchString(namePattern, name) // _ is used to ignore the boolean result of the regexp.MatchString function (which checks if name matches the namePattern regular expression).
	emailMatching, _ := regexp.MatchString(emailPattern, email)
	phoneMatching, _ := regexp.MatchString(phonePattern, phone)

	//printing output

	fmt.Printf("Valid Name=  %t\n", nameMatching) // When you use %t, it expects an argument of type bool.
	fmt.Printf("Valid Email= %t\n", emailMatching)
	fmt.Printf("Valid Phone= %t\n", phoneMatching)

}
