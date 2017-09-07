package main

import (
	"fmt"
)

func main() {
	fmt.Println(Lcm(10, 15))
}

//Lcm calculates the lowest common multiple using euclids algorithm
func Lcm(a, b int) int {
	return (a * b) / Gcd(a, b) //the Lowest Common Multiple is product of a and b
	//divided by the greatest common denominator
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
