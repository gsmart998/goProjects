package main

import "fmt"

type Human struct {
	name string
	sex  string
	age  int
}

type Employee struct {
	Human
	position string
	salary   int
}

func main() {
	fmt.Println("Hello, World!")
	h1 := Human{"Ivan", "Male", 23}
	h2 := Human{"Yana", "Female", 39}
	e1 := Employee{Human{"Dima", "Male", 30}, "Manager", 20000}
	e2 := Employee{h1, "Engineer", 40000}
	fmt.Println(h1, h2, e1, e2)
	fmt.Println(h1.age, e2.age, e1.name, e2.salary)
}
