package main

import "fmt"

func TypedocFunc() {
	type NoKTP string

	var KtpAsraf NoKTP = "888888888"
	fmt.Println(KtpAsraf)
	fmt.Println(NoKTP("121212121212"))
}
