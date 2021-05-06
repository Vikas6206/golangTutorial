package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

// issues without the use of interfaces, similar functionality but 2 functions
// methos overloading is not supported in go language

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	// custome logic to get the message in english
	return "Hi there"
}

func (sb spanishBot) getGreeting() string {
	// custome logic to get the message in spanish
	return "Hola"
}
