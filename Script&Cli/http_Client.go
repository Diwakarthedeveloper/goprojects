// Declare the package name. All Go programs start running in package main.
package main

// Import necessary packages
import (
	"bytes"    // Implements functions for the manipulation of byte slices.
	"fmt"      // Implements formatted I/O with functions analogous to C's printf and scanf.
	"io"       // Provides basic interfaces to I/O primitives.
	"net/http" // Provides HTTP client and server implementations.
)

// The main function, where the program starts.
func main() {
	// Define the base URL for httpbin
	mypage := "https://httpbin.org"

	// Perform an HTTP GET request
	getResp, galtiya := http.Get(mypage + "/get")
	// If there is an error (galtiya is not nil), print the error and return from the function
	if galtiya != nil {
		fmt.Println("Error on GET request:", galtiya)
		return
	}
	// Ensure the response body is closed when the function returns
	defer getResp.Body.Close()
	// Read all data from the response body
	getBody, _ := io.ReadAll(getResp.Body)
	// Print the GET response
	fmt.Println("GET Response:", string(getBody))

	// Define the body for the HTTP POST request
	postBody := []byte(`{"key": "value"}`)
	// Perform an HTTP POST request
	postResp, galtiya := http.Post(mypage+"/post", "application/json", bytes.NewBuffer(postBody))
	// If there is an error, print the error and return from the function
	if galtiya != nil {
		fmt.Println("Error on POST request:", galtiya)
		return
	}
	// Ensure the response body is closed when the function returns
	defer postResp.Body.Close()
	// Read all data from the response body
	postBodyResp, _ := io.ReadAll(postResp.Body)
	// Print the POST response
	fmt.Println("POST Response:", string(postBodyResp))

	// Create a new HTTP client
	client := &http.Client{}
	// Define the body for the HTTP PUT request
	putBody := []byte(`{"key": "updated value"}`)
	// Create a new HTTP PUT request
	putReq, galtiya := http.NewRequest(http.MethodPut, mypage+"/put", bytes.NewBuffer(putBody))
	// Set the Content-Type header
	putReq.Header.Set("Content-Type", "application/json")
	// Send the HTTP PUT request
	putResp, galtiya := client.Do(putReq)
	// If there is an error, print the error and return from the function
	if galtiya != nil {
		fmt.Println("Error on PUT request:", galtiya)
		return
	}
	// Ensure the response body is closed when the function returns
	defer putResp.Body.Close()
	// Read all data from the response body
	putBodyResp, _ := io.ReadAll(putResp.Body)
	// Print the PUT response
	fmt.Println("PUT Response:", string(putBodyResp))

	// Create a new HTTP DELETE request
	deleteReq, galtiya := http.NewRequest(http.MethodDelete, mypage+"/delete", nil)
	// Send the HTTP DELETE request
	deleteResp, galtiya := client.Do(deleteReq)
	// If there is an error, print the error and return from the function
	if galtiya != nil {
		fmt.Println("Error on DELETE request:", galtiya)
		return
	}
	// Ensure the response body is closed when the function returns
	defer deleteResp.Body.Close()
	// Read all data from the response body
	deleteBodyResp, _ := io.ReadAll(deleteResp.Body)
	// Print the DELETE response
	fmt.Println("DELETE Response:", string(deleteBodyResp))
}
