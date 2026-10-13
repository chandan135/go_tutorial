package main

import "fmt"

func main() {
	fmt.Println("---Slices---")

	names := []string{"Alice", "Bob", "John"}
	fmt.Println(names)

	items := make([]int, 3, 10)
	fmt.Printf("items: %v\n", items)
	fmt.Printf("length: %d\n", len(items))
	fmt.Printf("capacity: %d\n", cap(items))

	fmt.Println("---Testing Slices---")

	for i := 0; i < 14; i++ {
		items = append(items, i)
	}
	for index, item := range items {
		fmt.Printf("index= %d && item=%d\n", index, item)
	}
	fmt.Printf("length: %d\n", len(items))
	fmt.Printf("capacity: %d\n", cap(items))

	fmt.Println("Slicing from 10th to 15th element -->", items[10:16])
}

// In Go , Slices are dynamic arrays
// The zero value of slice is nil
// Slices are having length and capacity
