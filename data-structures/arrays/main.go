package main

import "fmt"

func main() {

	var numbers [2]int
	fmt.Printf("%+v\n", numbers)
	numbers[0] = 5
	numbers[1] = 10
	fmt.Printf("element-1: %d\n", numbers[0])
	fmt.Printf("element-2: %d\n", numbers[1])

	primeNumbers := [4]int{2, 3, 5, 7}
	fmt.Printf("%d\n", primeNumbers)

	for i := 0; i < len(primeNumbers); i++ {
		fmt.Printf("%d\n", primeNumbers[i])
	}

	var matrix [2][3]int
	var sum int = 0
	fmt.Printf("%d\n", matrix)
	fmt.Println("===Matrix Operation Performed===")
	for i := 0; i < 2; i++ {
		for k := 0; k < 3; k++ {
			sum = i + k
			matrix[i][k] = sum
		}
		sum = 0
	}

	fmt.Printf("%d\n", matrix)

}
