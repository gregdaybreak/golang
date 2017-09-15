package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Factor(204))
}

//Factor finds the factors of an integer and returns a slice of the factors
//O(√2^n)
func Factor(a int) []int {
	factors := make([]int, 0, 6) //make a new slice to hold the factors
	for a%2 == 0 {               //lets check the 2's first to remove them
		factors = append(factors, 2) //append 2 to the slice if 2 is a factor
		a = a / 2                    //divide the number by 2 to remove the 2's
	}
	factor := 3                     //start at things that could be a factor of 3
	stopAt := math.Sqrt(float64(a)) //we will stop when the factor is less than
	//the square root of the number
	for factor <= int(stopAt) {
		//pull out multiples of factor
		for a%factor == 0 {
			//pull out this factor
			factors = append(factors, factor)
			a = a / factor
			//calculate new stop_at value
			stopAt = math.Sqrt(float64(a))
		}
		//go to the next factor
		factor += 2

	}
	if a > 1 { //if we have a remaining number that cannot be factored add this
		//number to the slice
		factors = append(factors, a)
	}
	return factors //return the slice
}
