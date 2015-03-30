//https://projecteuler.net/problem=1
//If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

//Find the sum of all the multiples of 3 or 5 below 1000.

package main

import (
    "fmt"
    "strconv"
)

func main() {
    sum := findSum();
    fmt.Println("Sum of all the multiples of 3 or 5 below 1000 is "+ strconv.Itoa(sum)); //Sum of all the multiples of 3 or 5 below 1000 is 234168
}

func findSum() int{
    sum := 0;
    for i:=0;i<1000;i++{
        if(i%3==0 || i%5==0){
            sum = sum+i;
        }
    }
return sum
}