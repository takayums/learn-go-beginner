package main

import "fmt"

func main() {
	name := "Eko"

	if name == "Eko" {
		fmt.Println("Betul")
	} else {
		fmt.Println("Salah")
	}

	if length := len(name); length < 4 {
		fmt.Println("Kurang dari 4")
	} else {
		fmt.Println("Lebih dari 4")
	}
}
