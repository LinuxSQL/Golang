package main

import (
	"fmt"
)

type user struct {
	name string
	age  byte
}

type manager struct {
	user
	title string
}

func (u user) ToString() string {
	return fmt.Sprintf("%+v", u)
}

func main() {
	var m manager
	m.name = "SQL"
	m.age = 22
	m.title = "CTO"

	println(m.ToString())
}
