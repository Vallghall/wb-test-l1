package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewFloat(12345678) // init big.Float
	b := big.NewFloat(56473892) // init big.Float

	sum := new(big.Float).Add(a, b) // method for summing big.Floats
	sub := new(big.Float).Sub(a, b) // method fot subtracting big.Floats
	rem := new(big.Float).Quo(a, b) // method for dividing big.Floats
	mul := new(big.Float).Mul(a, b) // method for multiplying big.Float

	// print the above results out
	fmt.Printf(
		"Sum: %1.f\nSub: %1.f\nRem: %1.3f\nMul: %1.f\n",
		sum,
		sub,
		rem,
		mul,
	)
}
