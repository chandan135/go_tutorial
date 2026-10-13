package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println("---Slice Advanced---")
	original := []int{0, 1}
	for i := 2; i < 10; i++ {
		original = append(original, i)
	}
	fmt.Println(original)
	fmt.Printf("original: %v, len: %d, cap:%d\n", original, len(original), cap(original))
	original = slices.Insert(original, 1, 1000)
	fmt.Println(original)

	s1 := original[1:7]
	for i, value := range s1 {
		fmt.Printf("address is %v and value is %d\n", &i, value)
	}

	fmt.Println("---checking the address of an array elements---")

	var exampleArray = [4]int{1, 2, 3, 4}
	for i, value := range exampleArray {
		fmt.Printf("address is %v and value is %d\n", &i, value)
	}
}
