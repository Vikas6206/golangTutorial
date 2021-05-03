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
		contact: contactInfo{
			email:   "vkr@gmailt.com",
			zipCode: 9400,
		},
	}
	//go is pass by value i.e herw vikas4 struct is at one mermory location
	// and when go calls updaeName all the infor from vikas4 is copied to
	// p person which is at different address in the memory and hence the value is
	// not updated in the vikas4 struct

	vikas4.updateName("Vivek")
	vikas4.print()
	//this can be resolved with usin pointer, pointer to person
	// vikasPointer := &vikas4
	// vikasPointer.updateNameBypointer("Vivek")
	//we can even use the shortcut
	vikas4.updateNameBypointer("Vivek")
	vikas4.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
	fmt.Println()
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName

}

// this functiona accept a type that is pointerto a person
func (pointerToPerson *person) updateNameBypointer(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName

}
