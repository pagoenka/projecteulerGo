//https://projecteuler.net/problem=1
//If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

//Find the sum of all the multiples of 3 or 5 below 1000.

//Another way of solving the problem 1
package main

import (
    "fmt"
    "strconv"
)

func main() {
    sum := findSumV2();
    fmt.Println("Sum of all the multiples of 3 or 5 below 1000 is "+ strconv.Itoa(sum)); //Sum of all the multiples of 3 or 5 below 1000 is 233168
}

func findSumV2() int{
    sum3multiliers := 333*(3+999)/2;
    sum5multiliers := 200*(5+995)/2;
    sum15multiliers := 66*(15+990)/2;

    sum:= sum3multiliers + sum5multiliers - sum15multiliers;
    return sum
}