package main

import "fmt"

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

// printGreeting() contains the same logic but different parameters
func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}

func (eb englishBot) getGreeting() string {
	// custome logic to get the message in english
	return "Hi there"
}

func (sb spanishBot) getGreeting() string {
	// custome logic to get the message in spanish
	return "Hola"
}
