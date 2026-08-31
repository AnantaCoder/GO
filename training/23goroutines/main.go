package main

import (
	"fmt"
	"net/http"
	"time"
)

func greeter(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go greeter("Hello")
	greeter("world")



	var urls = []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.ai.com",
		
	}
	for _,url := range urls{
		status,err := getStatusCode(url)
		if err != nil {
		fmt.Println("There is an error : ", err)
		return
	    }

		fmt.Println(fmt.Sprintf("Status code for %s is %d", url, status))


	}

}

// we can use async from sync

func getStatusCode(endpoint string) (int, error) {

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("There is an error : ", err)
		return 0, err
	}
	defer res.Body.Close()
	return res.StatusCode, nil
}
