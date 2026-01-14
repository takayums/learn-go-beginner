package main

import "fmt"

func TipeDataFunc() {
	// 1. Tipe Data Numerik Non Desimal
	// Berisi nilai bilangan bulat (bilangan positif dan bilangan negatif) dan juga bilangan cacah (bilangan positif)
	// bilangan cacah "uint"
	// bilangan bulat "int"
	var positiveNumber uint8 = 89
	negativeNumber := -1243423644

	fmt.Printf("bilangan positif %d \n", positiveNumber)
	fmt.Printf("bilangan negative %d \n", negativeNumber)

	// 2. Tipe Data Numberik Desimal
	decimalNumber := 2.35
	fmt.Printf("bilangan decimal %f \n", decimalNumber)

	// 3. Tipe Data Boolean
	// Tipe data boolean yaitu hanya true & false
	var exits bool = true
	fmt.Printf("Nilai boolean benar yaitu %t \n", exits)

	// 4. TIpe Data String
	// Tipe data yang terdiri dari karakter dan diapit oleh ""
	var message string = "Hallo boy"
	fmt.Printf("message is %s \n", message)

	// 5. Nilai nil dan Zero Value
	// Niali nil yaitu bukan tipe data melainkan sebuah nilai kosong.
	// Zero value yaitu nilai default dari sebuah tipe data jika tipe data tersebut tidak bernilai
	// boolean = false
	// string = ""
	// numberik nondecimal = 0
	// numerik decimal =  0.0

	var nilai32 int32 = 129
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)
	fmt.Println(nilai8)

	// Konversi string
	name := "Asraf"
	e := name[0]
	eString := string(e)

	fmt.Println(eString)
}
