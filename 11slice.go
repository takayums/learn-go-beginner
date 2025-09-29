package main

import "fmt"

func main() {
	names := [...]string{"Asraf", "As'ad", "Afi"}
	slice := names[1:3]

	fmt.Println(slice[0])
	fmt.Println(slice[1])

	// Menambahkan / Append Slice
	days := [...]string{"senin", "selasa", "rabo", "kamis", "jum'at", "sabtu", "minggu"}
	daySlice1 := days[5:]
	fmt.Println(daySlice1)

	daySlice1[0] = "Sabu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(days) // days + Sabtu Baru + Minggu Baru

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Ups"
	fmt.Println(daySlice2)
	fmt.Println(days)

	// Membuat Slice
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Eko"
	newSlice[1] = "Eko"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
}
