package main

import "fmt"

func main() {
	fmt.Println(Sieve(100))

}

//Sieve lists prime numbers up to the limit using the Sieve of Eratosthenes
func Sieve(limit int) []int {
	sieve := make([]bool, limit)    //new slice of bool (set to false) with length of limit
	primes := []int{2}              //new slice of int with 2 already in it (it always has 2)
	for i := 3; i < limit; i += 2 { //while i is less than the limit and increment by 2 to get rid of
		//multiples of 3, 5, 7, etc.
		if sieve[i] == false { //if the value in the array is false
			for p := i * i; p < limit; p += i { //p equals i squared (so 9 the first time)
				//after the loop it adds i to p (so 12 the first time 9+3)
				sieve[p] = true //mark that spot in the array as true
			}
			primes = append(primes, i) //otherwise add the number to the primes slice
		}
	}
	return primes
}
