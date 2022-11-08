package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)

	var mark person

	mark.print()

	mark.firstName = "Mark"
	mark.lastName = "Zuckeberg"

	mark.print()

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 89012510,
		},
	}

	jim.print()

	jim.updateName("Jimmy")
	jim.print()

	jimPointer := &jim
	fmt.Println(jimPointer)
	jimPointer.updateName("McGill")
	jimPointer.print()
}

func (p person) print() {
	fmt.Printf("%+v \n", p)
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
