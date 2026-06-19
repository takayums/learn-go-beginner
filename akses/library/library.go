package library

import "fmt"

// public exported
func SayHello(name string) {
	fmt.Println("hello gaes")
	introduce(name)
}

// prive / unexported
func introduce(name string) {
	fmt.Println("Namaku", name)
}
