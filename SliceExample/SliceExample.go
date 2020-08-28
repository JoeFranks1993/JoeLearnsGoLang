package main

import "fmt"

func main() {
	// Create Slice.
	var list []int // Or s := make([]string, 3) to create it with a default length.

	// Append an element to the slice.
	list = append(list, 10)
	list = append(list, 3)
	list = append(list, 21)
	list = append(list, 67)

	// Append multiple items to slice with one invokation.
	list = append(list, 21, 80, 56, 21)

	// Loop through the slice and print each elements.
	for i := 0; i < len(list); i++ {
		fmt.Println(list[i])
	}
}
