package main

import "fmt"

func ArrayFunc() {
	var names [3]string

	names[0] = "Asraf"
	names[1] = "Budi"
	names[2] = "As'ad"

	fmt.Println(names)

	// Membuat Array Langsung
	array1 := [3]int{1, 2, 3}
	fmt.Println(array1)

	// Membuat Array Tanpa Jumlah-nya
	array2 := [...]int{
		90, 80, 70, 60,
	}
	fmt.Println(array2)
	array2[0] = 100
	fmt.Println(array2)
}
