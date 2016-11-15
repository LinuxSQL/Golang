package main

import "fmt"

type user struct {
	name string
	age  byte
}

func (u user) Print() {
	fmt.Printf("%+v\n", u)
}

// interface type
type Printer interface {
	Print()
}

func main() {
	var u user
	u.name = "SQL"
	u.age = 22

	var p Printer = u
	p.Print()
}
