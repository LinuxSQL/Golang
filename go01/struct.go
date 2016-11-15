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

func main() {
	var m manager

	m.name = "SQL"
	m.age = 22
	m.title = "CTO"

	fmt.Println(m)
}
