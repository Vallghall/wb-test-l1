package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "abCsdfghAjkl"       // init str
	fmt.Println(allUnique(a)) // print out if unique or not
}

func allUnique(s string) bool {
	// init map
	m := make(map[rune]struct{})
	// range over lower-cased string for case-insensitive check
	for _, val := range strings.ToLower(s) {
		// if map doesn't have a value, i.e. the value is
		// first-met, continue ranging. Otherwise, return false
		if _, ok := m[val]; !ok {
			m[val] = struct{}{}
			continue
		}
		return false
	}
	return true
}
