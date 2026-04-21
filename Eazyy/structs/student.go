package structs

import "fmt"

type Student struct {
	name  string
	grade int
}

func Func_student() {
	s1 := Student{name: "Abhiram  A R", grade: 10}
	s2 := Student{name: "Ananthu", grade: 12}

	fmt.Printf("%v is studying at %v grade\n", s1.name, s1.grade)
	fmt.Printf("%v is studying at %v grade\n", s2.name, s2.grade)
}
