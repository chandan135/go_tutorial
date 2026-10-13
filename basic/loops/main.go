package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("-----Control Flow and Logic-----")

	fmt.Println("---C style loop---")
	for i := 1; i <= 10; i++ {
		fmt.Println(i)

	}

	fmt.Println("---While Style Loop---")
	var k int = 3
	for k > 0 {
		fmt.Println(k)
		k--
	}

	fmt.Println("---Infinite Loop---")
	var counter int = 0
	for {
		fmt.Println("counter=", counter)
		counter++
		if counter >= 5 {
			break
		}
	}

	fmt.Println("---array---")
	items := [3]string{"Go", "Python", "c++"}
	for index, value := range items {
		fmt.Println(index, value)
	}

	// map
	fmt.Println("========Map=======")
	userAccess := map[string]bool{
		"Bob":   true,
		"Alice": true,
	}

	if hasAccess, ok := userAccess["Alice"]; ok && hasAccess {
		fmt.Println("user can access the system")
	} else {
		fmt.Println("Access Denied!!!")
	}

	// value, ok := mapName[key]
	// value contains the value associated with the key.
	// ok is a Boolean indicating whether the key exists.

	//Switch-case
	fmt.Println("---Switch-Case---")
	switch hour := time.Now().Hour(); {
	case hour < 12:
		fmt.Println("Good Morning")
	case hour < 17:
		fmt.Println("Good Afternoon")
	case hour < 21:
		fmt.Println("Good Evening")
	default:
		fmt.Println("Good Night")
	}

}
