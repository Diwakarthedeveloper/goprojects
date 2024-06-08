package main

import (
	"fmt"
	"os"
)

func main() {

	//below setting the enviroment variables
	os.Setenv("DJ_ENV_VARS", "Dello DJ")

	// reading the enviroment variables
	djvar := os.Getenv("DJ_ENV_VARS")
	if djvar == "" {
		fmt.Println("DJ_ENV_VARS is not set ")

	} else {
		fmt.Println("DJ_ENV_VARS:%s\n", djvar)
	}

	// printig down list of enviriment variables
	fmt.Println("List of all Enviroment cariables")
	for _, env := range os.Environ() {
		fmt.Println(env)
	}

}
