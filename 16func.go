package main

import "fmt"

// Func
func sayHello() {
	fmt.Println("Hello")
}

// Func Parameter
func sayHelloTo(name string, name2 string) {
	fmt.Println("Hello", name2, "Aku", name)
}

// Return Value
func getName(name string) string {
	return name
}

// Multiple Return Value
func getFullName() (string, string) {
	return "Asraf", "Eko"
}

// Memberikan nama return value
func getCompleteNmae() (fname, lname, mname string) {
	fname = "Eko"
	lname = "Lukman"
	mname = "Hadi"
	return fname, mname, lname
}

// Variadic Func
func sum(number ...int) int {
	total := 0
	for _, num := range number {
		total += num
	}
	return total
}

// Func Type Declaration
type Filter func(string) string

// Func as Parameter
func sayHelloWithFilter(name string, filter func(string) string) {
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func sayHelloWithFilterr(name string, filter Filter) {
	fmt.Println("Hello", filter(name))
}

// Func Anonymouse
type Blacklis func(string) bool

func registerUser(name string, blacklist Blacklis) {
	if blacklist(name) {
		fmt.Println("Your are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

// Func Recursive
func faktorial(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * faktorial(value-1)
	}
}

func main() {
	sayHello()

	sayHelloTo("asraf", "doni")

	nameAdit := getName("adit")
	fmt.Println(nameAdit)

	// Menghiraukan return value
	nameA, _ := getFullName()
	fmt.Println(nameA)

	fname, lname, mname := getCompleteNmae()
	fmt.Println(fname, lname, mname)

	total := sum(10, 10, 10, 10)
	fmt.Println("Total", total)

	numbers := []int{10, 101, 10}
	total2 := sum(numbers...)
	fmt.Println("Total 2", total2)

	// Func Value
	sayhello := sayHello
	sayhello()

	sayHelloWithFilter("Asraf", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)

	sayHelloWithFilterr("Okee", spamFilter)

	// Func Anonymouse
	blacklist := func(name string) bool {
		return name == "Anjing"
	}

	registerUser("eko", blacklist)
	registerUser("Anjing", blacklist)

	// Recursive Func
	nilaiFaktorial := faktorial(3)
	fmt.Println("faktorial", nilaiFaktorial)
}
