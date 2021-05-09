package main

import (
	"fmt"
	"net/http"
)

func main() {
	//check the staus of the web site to see if the website is up and running
	//this needs to be checked multiple times in a day
	links := []string{
		"http://www.google.com",
		"http://www.facebook.com",
		"http://www.stackoverflow.com",
		"http://www.golang.org",
		"http://www.amazon.com",
	}

	for _, link := range links {
		checkLink(link)
	}
}

//not returning any value
func checkLink(link string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}

	fmt.Println(link, "is up!")
}

//Problem: This prohram takes lots of time since its sequential i.e
// it executes one program get request waits for response and then validate and print
//it onto the console

//We can make it faster by executing the get request in parallel
