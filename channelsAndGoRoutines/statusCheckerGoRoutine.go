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

	//create channel which will be sharing a data of type string
	//from main routine to child routine or vice versa
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	//application is waiting for the data to come on the channel
	// once it gets the response it terminates the application

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}

}

//not returning any value
//channel which shares the data of type string
func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "Might be down I think"
		return
	}

	fmt.Println(link, "is up!")
	c <- "Yup its up"
}
