package slices

import "fmt"

func Append_log() {
	s := []int{}
	for i := 1; i < 5; i++ {
		s = append(s, i)
	}
	fmt.Println(s)
}
