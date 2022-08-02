package main

import "math"

func catsAndMouse(x int32, y int32, z int32) string {
	if x > 100 || y > 100 || z > 100 {
		panic("error")
	}
	// We need calculate the absolute distance between all points
	dist := math.Abs(float64(x - z))
	dist2 := math.Abs(float64(y - z))
	if dist < dist2 {
		return "Cat A"
	} else if dist > dist2 {
		return "Cat B"
	} else {
		return "Mouse C"
	}
}

func main() {
	catsAndMouse(1, 2, 3)
}
