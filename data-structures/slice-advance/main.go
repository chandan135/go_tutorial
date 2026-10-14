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

	var newSlice []int
	fmt.Println()
	fmt.Println("---Examining the capacity and length of the new slice---")

	fmt.Printf("\ncapacity = %d\nlength = %d\n", cap(newSlice), len(newSlice))
	fmt.Println(newSlice)

	for i := 0; i < 5; i++ {
		newSlice = append(newSlice, i)
	}
	fmt.Printf("\ncapacity = %d\nlength = %d\n", cap(newSlice), len(newSlice))
	fmt.Println(newSlice)

	for j := 5; j < 10; j++ {
		newSlice = append(newSlice, j)
	}
	fmt.Printf("\ncapacity = %d\nlength = %d\n", cap(newSlice), len(newSlice))
	fmt.Println(newSlice)

	for k := 10; k < 15; k++ {
		newSlice = append(newSlice, k)
	}
	fmt.Printf("\ncapacity = %d\nlength = %d\n", cap(newSlice), len(newSlice))
	fmt.Println(newSlice)
}
