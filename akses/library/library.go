package library

import "fmt"

// besar kecil huruf dari struct baik nama type struct atau propertinya berpengaruh ketika berhubungan dengan exported dan unexported
type Student struct {
	Name  string
	Grade int
}

// public exported
func SayHello(name string) {
	fmt.Println("hello gaes")
	introduce(name)
}

// prive / unexported
func introduce(name string) {
	fmt.Println("Namaku", name)
}
