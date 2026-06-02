package main

import "fmt"

func SwitchFunc() {
	name := "Asraf"

	switch name {
	case "Asraf":
		fmt.Println("Asraf tenan")
	case "Eko":
		fmt.Println("Eko tenan")
	default:
		fmt.Println("Maaf tidak kenal")
	}

	switch length := len(name); length <= 5 {
	case true:
		fmt.Println("Name Benar")
	case false:
		fmt.Println("Nama terlalu panjang")
	}
}
