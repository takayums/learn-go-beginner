package main

import "fmt"

func main() {
	var a = 10
	var b = 20
	var c = a + b

	fmt.Println(c)

	// Augmented Sigments
	var i = 10
	i += 10
	fmt.Println(i)

	// Unary Opertor
	var j = i
	j++
	j++
	fmt.Println(j)

	// Operasi perbandingan
	var name1 = "Eko"
	var name2 = "Eko"

	var result bool = name1 == name2
	fmt.Println(result)

	// Operasi Boolean
	var nilaiakhir = 90
	var absensi = 80

	var statusLulusNilai bool = nilaiakhir > 80
	var statusLulusAbsensi bool = absensi >= 80

	var statusLulus = statusLulusAbsensi && statusLulusNilai
	fmt.Println(statusLulus)
}
