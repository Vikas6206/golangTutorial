package main

import "fmt"

func main() {

	// in slice , the vlues of the slice is modifed i.e here
	// it works as pass by reference evn though we don't use pointers here
	// but please note in case of struct we have to use pointers
	//this is the go gotcha that it functions differently that structs

	slice := []string{"hi", "whats", "up", "bruh?"}

	slice[3] = "nibba?"
	fmt.Println(slice)
}
