package main

import (
	"fmt"
	"strings"
)

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func createHugeString(n int) string {
	sb := strings.Builder{}
	for i := 0; i < n; i++ {
		sb.WriteString("А") // test with 2-byte symbols
	}
	return sb.String()
}

func main() {
	someFunc()                              // run func
	l := len(strings.Split(justString, "")) // get the actual number of characters
	fmt.Println(l)                          // though len is 100, the is 50 symbols

	// FIX

	fixed := []rune(createHugeString(1 << 10))
	fixed = fixed[:100]
	fmt.Println(string(fixed)) // now it always cuts to consistent number of characters
}
