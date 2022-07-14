package main

import "fmt"

func main() {
	str := "главрыба"         // init str
	fmt.Println(reverse(str)) // reverse it
}

func reverse(s string) string {
	rs := []rune(s) // assert to []rune type to take encoding into account

	// swap the i-th and the j-th values until j >= i
	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}

	// assert back to string type
	return string(rs)
}
