package main

import "golang.org/x/tour/pic"
//import "math"

func Pic(dx, dy int) [][]uint8 {
	arr := make([][]uint8, dy)
	
	for y := range arr{
		arr[y] = make([]uint8, dx)
		for x := range arr[y]{
			//arr[y][x] = uint8(math.Pow(float64(x), float64(y)))
			//arr[y][x] = uint8(x*y)
			arr[y][x] = uint8(x+y/2)
		}
	}
	return arr
}

func main() {
	pic.Show(Pic)
}