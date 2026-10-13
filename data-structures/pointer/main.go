package main

import "fmt"

func callByValue(val int) {
	val = val * 10
	fmt.Printf("Value of val is %d\n", val)
	fmt.Printf("Address of val is %v\n", &val)

}

func callByPointer(val *int) {
	*val = *val * 10
	fmt.Printf("Value of val is %d\n", *val)
	fmt.Printf("Address of val is %v\n", val)

}

func main() {
	fmt.Println("---Pointers---")
	var age int = 10
	agePtr := &age
	fmt.Printf("Memory address of age is %v\n", agePtr)
	fmt.Printf("Value of age is: %d\n", age)

	var refAgePtr **int = &agePtr
	fmt.Printf("Address of agePtr is:%v\n", &agePtr)
	fmt.Printf("Value of refAgePtr is:%v\n", refAgePtr)
	fmt.Println()

	var value int = 10
	fmt.Printf("Address of value is %v\n", &value)
	fmt.Printf("Value is %d\n", value)
	fmt.Println("---callByValue() Function call is being made---")
	callByValue(value)
	fmt.Println("---callByValue() Function call is done---")
	fmt.Printf("Address of value is %v", &value)
	fmt.Printf("Value is %d\n", value)

	fmt.Println("---callByPointer() Function call is being made---")
	callByPointer(&value)
	fmt.Println("---callByPointer() Function call is done---")
	fmt.Printf("Address of value is %v", &value)
	fmt.Printf("Value is %d\n", value)

}

// Pointer: It's a variable which stores address of another variable
