package main

import "fmt"

func VariableFunc() {
	// Menggunakan deklarasi tipe data
	var name string

	name = "Muhammad Asraf"
	fmt.Println(name)

	name = "Asraf"
	fmt.Println(name)

	// Tidak menggunakan deklarasi tipe data
	address := "Jalan Kalimantan 10"
	fmt.Println(address)

	// Deklarasi var dengan :=
	schools := "MAN 2 BANYUWANGI"
	fmt.Println(schools)

	// Deklarasi multi variable
	var (
		fname = "Asraf"
		lname = "Takayuma"
	)
	fmt.Println(fname)
	fmt.Println(lname)
}
