package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// create a slice of slices of uint8 with dy rows and dx columns.
	slice := make([][]uint8, dy)
	for i := range slice {
		// create a slice of uint8 with dx columns
		slice[i] = make([]uint8, dx)
		for j := range slice[i] {
			slice[i][j] = uint8((i + j) / 2)
		}
	}
	return slice
}

func main() {
	pic.Show(Pic)
}
