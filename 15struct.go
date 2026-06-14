package main

import (
	"fmt"
)

type student struct {
	name  string
	grade int
	age   int
}

type person struct {
	student
	age int
}

type animal struct {
	name string
	tipe string
}

func StructFunc() {
	var s1 student

	s1.name = "Asraf"
	s1.grade = 100

	fmt.Println("Name: ", s1.name)
	fmt.Println("Grade: ", s1.grade)

	s2 := student{"antok", 100, 24}
	fmt.Println("S2", s2)

	s3 := student{name: "askon"}
	fmt.Println("S3", s3)

	// struct Object Pointer
	s4 := student{name: "wicak", grade: 20}

	var s5 *student = &s4
	fmt.Println("S4", s4)
	fmt.Println("S5", s5)

	s5.name = "ganti"
	fmt.Println("S4", s4)
	fmt.Println("S5", s5)

	// embeded struct
	s7 := person{}
	s7.name = "R7"
	s7.age = 25
	s7.grade = 100

	fmt.Println("S7", s7)
	fmt.Println("S7 name", s7.name)
	fmt.Println("S7 name", s7.student.name)
	fmt.Println("S7 gade", s7.grade)
	fmt.Println("S7 gade", s7.student.grade)
	fmt.Println("S7 age", s7.age)

	// embeded property struct yang sama
	s7.student.age = 28
	fmt.Println("S7 age person", s7.age)
	fmt.Println("S7 age student", s7.student.age)

	// Pengisian nilai sub struct
	p1 := student{name: "asraf", age: 20}
	p2 := person{student: p1, age: 25}
	fmt.Println("Name", p1.name)
	fmt.Println("Name", p2.name)
	fmt.Println("Name", p2.age)
	fmt.Println("Name", p2.student.age)

	// anonymouse struct
	s10 := struct {
		animal
		color string
	}{
		animal: animal{"kucing", "omnivora"},
		color:  "blue",
	}

	fmt.Println("S10", s10)

	// kombinasi slice dan struct
	allAnimal := []animal{
		{"ayama", "omnivora"},
		{"bebek", "omnivora"},
		{"anjing", "karnivora"},
	}
	fmt.Println("All Slice from animal", allAnimal)

	// inisialisasi slice anonymouse struct
	allAnimal1 := []struct {
		animal
		color string
	}{
		{
			animal: animal{"ayama", "omnivora"},
			color:  "red",
		},
		{
			animal: animal{"bebek", "omnivora"},
			color:  "white",
		},
	}

	for _, animalItem := range allAnimal1 {
		fmt.Println(animalItem)
	}
	// deklarasi anonymouse struct menggunakan keyword var
	// nested struct
	// deklarasi dan inisialisasi struct secara horizontal
	// tag property dalam struct
	// type alias
}
