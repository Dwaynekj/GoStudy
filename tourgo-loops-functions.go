package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var z, znew float64 = 1.0, 1.0
	delta := 100.0

	for ; delta > 0.00004; {
		znew = z - ((z*z)-x)/(2*z)
		delta = math.Abs(float64(znew - z))
		z = znew
	}
	return z
}

func main() {
	fmt.Println(Sqrt(3))
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(1))
}
