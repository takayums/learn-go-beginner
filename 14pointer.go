package main

import "fmt"

func Pointer() {
	// var number *int
	// var name *string

	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("Number A (value): ", numberA)
	fmt.Println("Number A (address): ", numberB)

	fmt.Println("Number B (value): ", *numberB)
	fmt.Println("Number B (address): ", numberB)

	// variabel yang memiliki nilai pointer harus memiliki tipe data dengan tanda *tipedata ex. *int
	// Untuk akses nilai pointer pada sebuah variable dengan menggunakan & di awal variabel yang ingin dilihat pointernya
	// untuk melihat nilai dari sebuah variabel yang menyimpan pointer dengan cara *namavariable
}
