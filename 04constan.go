package main

import "fmt"

func ConsFunc() {
	// 1. Konstanta
	// Nilai yang sudah ada tidak dapat dirubah kembali

	const fname = "Asraf"
	const lname = "Takayuma"

	// error
	// fname = "yuma"
	// lname = "takayum"

	// 2. Deklarasi Multi Konstanta
	const (
		hname = "Asraf"
		tname = "Takayuma"
	)
	fmt.Printf("Nilai hname const: %s \n", hname)

	// error
	// hname = "yuma"
	// tname = "takayum"

	const satu, dua = 1, 2
	const three, four string = "tiga", "empat"
	fmt.Printf("Nilai satu: %d dua: %d \n", satu, dua)
}
