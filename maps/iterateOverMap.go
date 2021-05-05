package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "##ff0000",
		"green": "##wws00",
		"white": "##qo000",
	}
	fmt.Println("Original value :", colors)
	printMap(colors)
}

//func to accept the map and iterate over it

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println(key + " : " + value)
	}
}

//differnce between maps and structs
