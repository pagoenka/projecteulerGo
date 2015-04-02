//https://projecteuler.net/problem=4
//A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

//Find the largest palindrome made from the product of two 3-digit numbers.

//*** A palindrome is divisible by 11, so either of its factor should be divisible by 11 ***

package main

import (
    "fmt"
    "strconv"
)

func main() {
    maxPalinDrome :=0
    var product int;
    var subFactor int;
    for i:=999;i>=100;i--{
        var j int;
        if(i%11 == 0){
            j=i;
            subFactor = 1;
        }else{
            j = 990;
            subFactor = 11;
        }
        for j >= 100{
            product = i*j;
            if product > maxPalinDrome  {
                if isPalinDromeV2(product){
                    maxPalinDrome = product;
                }
            }else {
                break;
            }
            j = j - subFactor;
        }
    }
    fmt.Println("Largest palindrome made from the product of two 3-digit numbers is "+strconv.Itoa(maxPalinDrome)); //Largest palindrome made from the product of two 3-digit numbers is 906609

}

func isPalinDromeV2(n int) bool {
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
