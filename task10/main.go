package main

import (
	"fmt"
	"strings"
)

func main() {
	// init temperature data
	data := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// init map
	m := make(map[int][]float32)

	// range data
	for _, val := range data {
		cat := int(val) - (int(val) % 10) // get group category

		// init category group if not exists
		if _, ok := m[cat]; !ok {
			m[cat] = make([]float32, 0)
		}

		// add value to category group
		m[cat] = append(m[cat], val)
	}

	sb := strings.Builder{} // init strings.Builder
	for k, v := range m {

		// accumulate category group string
		for _, val := range v {
			sb.WriteString(fmt.Sprintf("%.1f, ", val))
		}
		fmt.Printf("%d: { %v }\n", k, sb.String())
		sb.Reset() // reset builder for reuse
	}
}
