package main

import "../utils"

func main() {
	var s []int
	utils.PrintSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	utils.PrintSlice(s)

	// the slice grows as needed.
	s = append(s, 1)
	utils.PrintSlice(s)

	// we can add more than one element at a time.
	s = append(s, 2, 3, 4)
	utils.PrintSlice(s)
}
