package main

import "fmt"

func main() {

	//different way to create a map

	//1st way
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#ff9000",
	}

	fmt.Println(colors)

	//2nd way - default value is assigned to it
	var colors1 map[string]string
	fmt.Println(colors1)

	//3rd way using make function
	colors3 := make(map[string]string)
	colors3["red"] = "#ff0000"
	fmt.Printf("%+v", colors3)

	//delete the key and value from map
	delete(colors3, "red")
	fmt.Printf(" After deletion ::::::::; %+v", colors3)

	//how to iterate over a map

}
