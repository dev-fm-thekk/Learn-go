package main

import (
	"Eazyy/arrays"
	"Eazyy/interfaces"
	"Eazyy/loops"
	"Eazyy/maps"
	"Eazyy/slices"
	"Eazyy/structs"
	"fmt"
)

func read_num() int {
	var num int
	fmt.Printf("Enter a number: ")
	fmt.Scan(&num)
	return num
}

func test_loops() {
	fmt.Println("Testing Loop: print_20.go")
	loops.Log()

	fmt.Println("Testing Loop: sum_n.go")
	result := loops.Sum_n(read_num())
	fmt.Printf("result: %v\n", result)

	fmt.Println("Testing Loop: mult_table.go")
	loops.Print_Mult_table(read_num())
}

func test_arrays() {
	fmt.Println("Testing Array: sum_array.go")
	result := arrays.Sum_arrays()
	fmt.Printf("Sum of array: %v\n", result)
	fmt.Println("Testing Array: max_array.go")
	input := [5]int{1, 2, 3, 4, 5}
	max := arrays.Max(input)
	fmt.Printf("Max of array: %v\n", max)
}

func test_slices() {
	fmt.Println("Test Slices: append_log.go")
	slices.Append_log()
	fmt.Println("Testing Slices: event_filter.go")
	even_slice := slices.Even_filter([]int{1, 2, 3, 4, 5})
	fmt.Println(even_slice)
}

func test_map() {
	fmt.Println("Test maps: word_fc.go")
	words := []string{"go", "is", "fun", "go", "rocks"}
	word_count := maps.Word_fc(words)
	fmt.Println(word_count)
	fmt.Println("Testing maps: country_map.go")
	maps.Country_map()
}

func test_interface() {
	fmt.Println("Test interface: shape.go")
	var s interfaces.Shape = interfaces.Circle{Radius: 5}
	fmt.Printf("Area of circle: %v\n", s.Area())
}

func test_structs() {
	fmt.Println("Test structs: student.go")
	structs.Func_student()

	fmt.Println("Test structs: Rectangle.go")

	rec := structs.Rectangle{Height: 12.0, Width: 16.0}
	fmt.Printf("Rectangle area: %v\n", rec.Area())
}
func main() {
	// test_loops()
	// test_arrays()
	// test_slices()
	// test_map()
	// test_interface()
	test_structs()
}
