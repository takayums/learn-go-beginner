package akses

import (
	"fmt"

	"github.com/takayums/learn-go-beginner/akses/library"
)

func LevelAkses() {
	library.SayHello("asraf")
	s1 := library.Student{"asraf", 20}
	fmt.Println("name", s1.Name)
	// fmt.Println("grade", s1.grade)
	// library.introduce("asraf")
}
