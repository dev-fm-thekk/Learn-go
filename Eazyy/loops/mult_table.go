package loops

import "fmt"

func Print_Mult_table(n int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%v * %v = %v\n", 5, i, 5*i)
	}
}
