package main

import (
	"fmt"
	"strings"
)

type Student struct {
	name string
	age  int
}

func (s Student) sayHello() {
	fmt.Println("Hello I'm Method Struct")
}

func (s Student) getName(i int) string {
	return strings.Split(s.name, " ")[i-1]
}

func (s Student) changeName(name string) {
	fmt.Println("Tanpa Pointer", name)
	s.name = name
}

// Pointer
func (s *Student) changeNamePointer(name string) {
	fmt.Println("Dengan Pointer", name)
	s.name = name
}

func Method() {
	s1 := Student{"Asraf Takayuma", 20}
	s1.sayHello()

	name := s1.getName(2)
	fmt.Println("Nama panggilan: ", name)

	s1.changeName("Muhammad")
	fmt.Println("Hasil Setelah Change Name", s1.name)

	// Pointer
	s1.changeNamePointer("Loha")
	fmt.Println("Hasil Setelah Change Name", s1.name)
}
