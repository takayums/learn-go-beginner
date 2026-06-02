package main

import "fmt"

func MapFunc() {
	person := map[string]string{
		"name": "Asraf", "address": "Jalan Banyuwangi",
	}

	fmt.Println("person", person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println("panjang map", len(person))
	person["name"] = "Anthina"
	fmt.Println(person["name"])
	delete(person, "name")
	fmt.Println("person", person)

	// Make Map
	book := make(map[string]string)
	book["title"] = "Skripsi"
	book["umur"] = "10"
	fmt.Println("book", book)
}
