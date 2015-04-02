//https://projecteuler.net/problem=4
//A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

//Find the largest palindrome made from the product of two 3-digit numbers.
package main

import (
    "fmt"
    "strconv"
)

func main() {
    maxPalinDrome :=0
    var product int;
    for i:=999;i>=100;i--{
        for j:=i;j>=100;j--{
            product = i*j;
            if product > maxPalinDrome  {
                if isPalinDrome(product){
                    maxPalinDrome = product;
                }
            }else {
                break;
            }

        }
    }
    fmt.Println("Largest palindrome made from the product of two 3-digit numbers is "+strconv.Itoa(maxPalinDrome)); //Largest palindrome made from the product of two 3-digit numbers is 906609

}

func isPalinDrome(n int) bool {
    originalNumber := n;
    rev := 0;

    for n > 0{
        rev = rev * 10 + (n%10);
        n = n/10;
    }
    if(rev == originalNumber){
        return true;
    }else {
        return  false;
    }
}
