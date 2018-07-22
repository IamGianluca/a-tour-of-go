package main

import "../utils"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	utils.PrintSlice(s)

	// slice the slice to give it zero length.
	s = s[:0]
	utils.PrintSlice(s)

	// extend its length.
	s = s[:4]
	utils.PrintSlice(s)

	// drop its first two values.
	s = s[2:]
	utils.PrintSlice(s)
}

