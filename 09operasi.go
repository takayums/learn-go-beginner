package main

import "fmt"

func OperasiFunc() {
	a := 10
	b := 20
	c := a + b

	fmt.Println(c)

	// Augmented Sigments
	i := 10
	i += 10
	fmt.Println(i)

	// Unary Opertor
	j := i
	j++
	j++
	fmt.Println(j)

	// Operasi perbandingan
	name1 := "Eko"
	name2 := "Eko"

	var result bool = name1 == name2
	fmt.Println(result)

	// Operasi Boolean
	nilaiakhir := 90
	absensi := 80

	var statusLulusNilai bool = nilaiakhir > 80
	var statusLulusAbsensi bool = absensi >= 80

	statusLulus := statusLulusAbsensi && statusLulusNilai
	fmt.Println(statusLulus)
}
