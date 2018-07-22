package main

import (
    "golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
    s := make([][]uint8, dx)
    for i, _ := range s {
        u := make([]uint8, dy)
		for idx, _ := range u {
			u[idx] = uint8(i) * uint8(idx)
		}
		s[i] = u
    }
    return s
}

func main() {
	pic.Show(Pic)
}
