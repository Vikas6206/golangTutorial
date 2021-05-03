package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	//different way to create a value of person
	//1 (depends on the orderof fields - drawback)
	vikas := person{"Vikas", "Kumar"}

	fmt.Println(vikas)
	//2 (with fieldname i.e. key value pair) better way so that we can add additional parameter

	vikas1 := person{
		firstName: "Vikas",
		lastName:  "Kumar"}

	fmt.Println(vikas1)

	//3 way
	var vikas3 person
	vikas3.firstName = "Vikas"
	vikas3.lastName = "Kumar"

	fmt.Println(vikas3)
	fmt.Printf("%+v", vikas3)

}
