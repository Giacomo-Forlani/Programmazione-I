package main

import "fmt"
import "math"

/*
 * Given an integer x, it returns true if x is prime
 */
func isPrime(x int) bool {
	for d := 2; float64(d) < math.Sqrt(float64(x)); d++ {
		if x % d == 0 {
			return false
		}
	}
	return true
}


/*
 * Given an integer x, it returns true iff x is an integer power of 2
 */
func isPowerOfTwo(x int) bool {
	for x > 1 {
		if x % 2 != 0 {
			return false
		}
		x /= 2
	}
	return true
}

/*
 * Given an integer x that is a power of two, returns the exponent of two y such that 2^y=x
 */
func powerOfTwoExponent(x int) int {
	c := 0
	for x > 1 {
		x /= 2
		c++
	}
	return c
}

/*
 * Given an integer x, it returns true iff x is a Mersenne prime
 */
func isMersennePrime(x int) bool {
	if !isPrime(x) {
		return false
	}
	if !isPowerOfTwo(x + 1) {
		return false
	}
	y := powerOfTwoExponent(x + 1)
	return isPrime(y)
}



func main() {
	for n := 2; ; n++ {
		if isMersennePrime(n) {
			fmt.Println(n)
		}
	}
}


