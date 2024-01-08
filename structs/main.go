package main

import "fmt"

type person struct {
	firsname string
	lastname string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	jim := person{
		firsname: "Maria",
		lastname: "Silva",
		contactInfo: contactInfo{
			email:   "mariajsilvat@gmail.com",
			zipCode: 121212,
		},
	}
	jimPointer := &jim
	jimPointer.updateName("Auri")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firsname = newFirstName
}

// Struct with receiver a function
func (p person) print() {
	fmt.Printf("%+v", p)
}
