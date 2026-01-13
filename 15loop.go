package main

import "fmt"

func LoopFunc() {
	for counter := 0; counter <= 5; counter++ {
		if counter == 4 {
			break
		}
		fmt.Println("counter-", counter)
	}

	names := []string{"Asraf", "Anthina", "Arkan"}
	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}
}
