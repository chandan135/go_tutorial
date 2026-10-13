package main

import "fmt"

// enum
// Go doesn't have enum keyord, its done using iota
const (
	sunday = iota
	monday
	tuesday
	wednesday
	thursday
	friday
	saturday
)

type LogLevel int

const (
	logError LogLevel = iota
	logWarn
	logInfo
	logDebug
)

func main() {
	fmt.Println(sunday)
	fmt.Println(monday)
	fmt.Println(tuesday)
	fmt.Println(wednesday)
	fmt.Println(thursday)
	fmt.Println(friday)
	fmt.Println(saturday)

	fmt.Println("Printing log related codes:")
	fmt.Println(logError)
	fmt.Println(logWarn)
	fmt.Println(logInfo)
	fmt.Println(logDebug)

}
