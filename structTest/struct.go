package main

import "fmt"

// Create new struct
type contact struct {
	firstName   string
	lastName    string
	phoneNumber string
	email       string
	addres      string
	dateOfBirth string
}

// Method show full struct info
func (c contact) printInfo() {
	fmt.Printf("Имя: %s\nФамилия: %s\nНомер телефона: %s\nE-Mail: %s\nАдрес: %s\nДата рождения: %s\n", c.firstName, c.lastName, c.phoneNumber, c.email, c.addres, c.dateOfBirth)
}

// Method for update struct
func (c *contact) setName(name string) {
	c.firstName = name
}

func main() {
	// Create new struct element
	c1 := contact{
		firstName:   "Вася",
		lastName:    "Пупкин",
		phoneNumber: "29384702934",
		email:       "asdfkl@gmail.com",
		addres:      "Moscow, Russia",
		dateOfBirth: "21.11.2002",
	}

	// Call method show full struct info
	c1.printInfo()

	// Call method to set new name
	c1.setName("Анатолий")

	// Call method show full struct info after update
	c1.printInfo()

}
