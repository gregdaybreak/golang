package main

import (
	"fmt"
)

func main() {
	fmt.Println(Gcd(78, 66))
}

//Gcd calculates the greatest common divisor of two numbers using Euclids
//algorithm
func Gcd(a, b int) int {
	for b != 0 { //loop runs as long as b is not equal to zero
		remainder := a % b //store the remainder of a divided by b
		//the GCD of a and b is the same as the GCD of B and the remainder (a % b)
		a = b         //a then equals b
		b = remainder //and b equals the remainder and repeat the loop
	}
	return a
}
