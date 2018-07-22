package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    calls := 0
    primes := []int{0, 1}
    return func() int {
        calls += 1
        if calls < 2 {
            return primes[calls]
        } else {
            new_prime := 0
            for _, v := range primes[len(primes)-2:] {
                new_prime += v
            }
            primes = append(primes, new_prime)
            return primes[len(primes)-1]
        }
    }
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
