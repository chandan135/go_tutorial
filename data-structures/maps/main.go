package main

import "fmt"

func main() {
	fmt.Println("---Maps---")

	studentGrades := map[string]int{
		"Alice": 90,
		"Bob":   90,
		"Jhon":  90,
	}

	fmt.Printf("%v\n", studentGrades)

	studentGrades["Bob"] = 92 // overriding value in a map
	fmt.Println(studentGrades)

	// To check whether a key exist in a map
	alice, ok := studentGrades["Alice"]
	fmt.Println(alice, ok)

	_, ok = studentGrades["Tesla"]
	fmt.Println(ok)

	if _, ok := studentGrades["Bob"]; ok {
		fmt.Println("Bob's marks is:", studentGrades["Bob"])
	}

	delete(studentGrades, "Alice") // deleting from a map
	fmt.Println(studentGrades)

	// Creating map from make keyword
	configs := make(map[string]int)
	fmt.Printf("%v %T\n", configs, configs)
	fmt.Printf("%v\n", configs == nil)

	// Declaring a map variable
	var exampleMap map[string]int
	fmt.Printf("exampleMap is a map of type %T\n", exampleMap)
	fmt.Printf("%v\n", exampleMap == nil)
}

// Key-value pair datastructure
// Zero value of a map is nil

// If map is created using make keyword, it wont be nil
// as memory initialization will take place
