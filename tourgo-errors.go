package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint(float64(e))
	//fmt.Sprint(e) will cause an infinite loop
	// Not sure why?
}

func Sqrt(x float64) (float64, error) {
	var e ErrNegativeSqrt = -2
	if x < 0 {
		return x, e
	
	}
	var z, znew float64 = 1.0, 1.0
	delta := 100.0

	for ; delta > 0.00004; {
		znew = z - ((z*z)-x)/(2*z)
		delta = math.Abs(float64(znew - z))
		z = znew
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}