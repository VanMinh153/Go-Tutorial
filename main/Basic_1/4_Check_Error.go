package main

import (
	"fmt"
	"strconv"
)

func convertToInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("couldn't convert %q to an integer:\n\t  %w", s, err)
	}
	return i, nil
}

func main() {
	inputs := []string{"42", "not a number", "2k3"}

	for _, input := range inputs {
		if i, err := convertToInt(input); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("Converted %q to %d\n", input, i)
		}
	}
}
