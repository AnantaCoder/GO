package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

const baseURL = "https://jsonplaceholder.typicode.com/posts"

func main() {
	// 1. Send POST request
	postBody := `{
		"title": "Test Title",
		"body": "This is a test post.",
		"userId": 1
	}`
	// base url content type , content type of request body,
	//  postbody is a variable which contains json data sent to the server
	// []]byte(postBody): This converts the postBody to a byte slice, as bytes.NewBuffer expects a byte slice as input.
	// bytes.NewBuffer(...): This creates a new bytes.Buffer from the byte slice, which implements the io.Reader interface required by http.Post.
	resp, err := http.Post(baseURL, "application/json", bytes.NewBuffer([]byte(postBody)))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("POST Response Status:", resp.Status)

	bodyBytes, _ := io.ReadAll(resp.Body)
	fmt.Println("POST Response Body:\n", string(bodyBytes))

	// 2. Send GET request to retrieve a specific post
	getURL := baseURL + "/1"
	
	getResp, err := http.Get(getURL)
	if err != nil {
		panic(err)
	}
	defer getResp.Body.Close()

	fmt.Println("\nGET Response Status:", getResp.Status)

	getBody, _ := io.ReadAll(getResp.Body)
	fmt.Println("GET Response Body:\n", string(getBody))
}
