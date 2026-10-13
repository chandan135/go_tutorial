package main

import "fmt"

const HOST = "localhsot"

func main() {
	fmt.Println("Hello World!!!")

	var interface_ex any
	var array = [3]int{1, 2, 3}
	fmt.Println(array)
	fmt.Println(interface_ex)

	var int_array = [5]int{}
	fmt.Println(int_array)

	var greeting string = "Hello Sam!!"
	fmt.Println(greeting[0:3], greeting[3:7])

	var fName, lName string
	fName = "Na"
	lName = "vy"
	fmt.Println(fName + lName)

	const PI = 3.22
	fmt.Printf("%v\n", PI)

	fmt.Println(HOST)

}

// go run, go build, go fmt, go get/install
