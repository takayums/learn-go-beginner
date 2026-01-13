package main

import "fmt"

func VariableFunc() {
	// 1. Menggunakan deklarasi tipe data dan juga menggunakan var
	var name string

	name = "Muhammad Asraf"
	fmt.Println(name)

	name = "Asraf"
	fmt.Println(name)

	// 2. Tidak menggunakan deklarasi tipe data
	// Atau disebut dengan type inference
	address := "Jalan Kalimantan 10"
	fmt.Println(address)

	// Deklarasi var dengan :=
	schools := "MAN 2 BANYUWANGI"
	fmt.Println(schools)

	// 3. Deklarasi multi variable
	var (
		fname = "Asraf"
		lname = "Takayuma"
	)
	fmt.Println(fname)
	fmt.Println(lname)

	one, two, tree := 1, 2, 3
	fmt.Printf("angka dari: %d %d %d\n", one, two, tree)

	// 4. Variable Underscore / tidak digunakan
	// Variable yang tidak digunakan tidak boleh dideklarasikan, dan harus diganti dengan "_"
	_ = "Hello Ini tidak digunakan"
	nameEmpty, _ := "Fani", "Andi"
	fmt.Printf("Nama Pertama %s\n", nameEmpty)

	// 5. Deklarasi dengan "new"
	nameFull := new(string)
	fmt.Println(nameFull)
	fmt.Println(*nameFull)

	// 6. Deklarasi variable dengan "make"
	// Berlaku pada channel, slice, dan map
}
