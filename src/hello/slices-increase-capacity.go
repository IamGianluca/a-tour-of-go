package main

import "../utils"

func increase_cap(s []int) []int {
    t := make([]int, len(s), (cap(s)+1)*2) // +1 in case cap(s) == 0
    copy(t, s)
    return t
}

func main() {
    s := []int{0, 1, 2}
    utils.PrintSlice(s)
    d := increase_cap(s)
    utils.PrintSlice(d)
}
