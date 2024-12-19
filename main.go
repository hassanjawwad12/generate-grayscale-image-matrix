package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	image := make([][]uint8, dy)

	// Loop through each row
	for y := 0; y < dy; y++ {
		// Create a slice for each row with length dx
		row := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			row[x] = uint8((x + y) / 2) 
		}
		image[y] = row // Assign the row to the outer slice
	}

	return image
}

func main() {
	pic.Show(Pic)
}
