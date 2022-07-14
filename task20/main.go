package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun fun — fun sun dog snow" // init str
	fmt.Println(reverseWords(str))               // reverse words
}

// reverseWords reverses the words in a given string
func reverseWords(s string) string {
	ss := strings.Split(s, " ") // split to words

	// swap the i-th and the j-th values until j >= i
	for i, j := 0, len(ss)-1; i < j; i, j = i+1, j-1 {
		ss[i], ss[j] = ss[j], ss[i]
	}

	// return a word-reversed string
	return strings.Join(ss, " ")
}
