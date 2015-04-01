//https://projecteuler.net/problem=3
//The prime factors of 13195 are 5, 7, 13 and 29.

//What is the largest prime factor of the number 600851475143 ?
package main

import (
    "fmt"
    "strconv"
)

func main() {
    sum := findPrimeFactor();
    fmt.Println("Largest prime factor is "+ strconv.Itoa(sum)); //Largest prime factor is 6857
}

func findPrimeFactor() int{
        n := 600851475143;
    maxPrime := 1;
    for n%2 == 0 {
        maxPrime = 2;
        n = n/2;
    }
    for i:=3;i <= n;i +=2 {
        for n%i == 0 {
            maxPrime = i;
            n = n/i;
        }
    }
    return maxPrime;
}