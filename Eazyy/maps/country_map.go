package maps

import "fmt"

func Country_map() {
	country_maps := map[string]string{
		"India":   "New Delhi",
		"USA":     "New York",
		"Germany": "Berlin",
	}

	for key, val := range country_maps {
		fmt.Printf("Capital of %v is %v\n", key, val)
	}
}
