package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
	Name string
	Age  int
}

func (dog Dog) Speak() string {
	return "Woof!"
}

func main() {
	dog := Dog{Name: "Buddy", Age: 3}
	var animal Animal = dog
	fmt.Println(animal.Speak())
}
