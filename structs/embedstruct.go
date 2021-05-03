package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {

	vikas4 := person{
		firstName: "Vikas",
		lastName:  "Kumar",
		//note that we need to provide the struct type contactInfo here
		//also note that all the line would have , since we are defining
		//multiline struct
		contact: contactInfo{
			email:   "vkr@gmailt.com",
			zipCode: 9400,
		},
	}
	vikas4.updateName("Vivek")
	vikas4.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName

}
