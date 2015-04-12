//https://projecteuler.net/problem=5

//2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

//What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
package main

import (
	"fmt"
	"strconv"
)

func main() {
	lcm(20)
}

func lcm(n int) {
	lcm := 1
	for i := 2; i <= n; i++ {
		gcd := gcd(lcm, i)
		lcm = (lcm / gcd) * i
	}

	fmt.Println("LCM is" + strconv.Itoa(lcm)) // LCM is 232792560
}

func gcd(n1 int, n2 int) int {
	for n2 != 0 {
		t := n2
		n2 = n1 % n2
		n1 = t
	}
	return n1
}
