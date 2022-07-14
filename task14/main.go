package main

import "fmt"

func main() {
	var foo any = "bar" // init value of type any (interface{})

	// type switch
	switch foo.(type) {
	case string:
		fmt.Println("foo is a string")
	case int:
		fmt.Println("foo is an int")
	case bool:
		fmt.Println("foo is a bool")
	case chan any:
		fmt.Println("foo is a chan")
	default:
		fmt.Println("foo is of some other type")
	}
}
