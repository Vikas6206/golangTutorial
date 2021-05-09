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

	//infininte loop
	// for {
	//watch channel c and whenever value comes out assign it to l
	for l := range c {
		//wait for the message to be received in the channel
		//after receiving the data in the channel call the go routine again
		// go checkLink(<-c, c)
		go checkLink(l, c)
	}

}

//not returning any value
//channel which shares the data of type string
func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		//to make it work like a status cheker i.e. whenever the website is down
		//return the link to the channel
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	//to make it work like a status cheker i.e. whenever the website is up
	//return the link to the channel
	c <- link
}
