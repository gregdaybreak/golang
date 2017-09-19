package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	argString, _ := strconv.Atoi(os.Args[1]) //convert argument to int
	myMap := map[bool]int{true: 0, false: 0} //new map to hold bool count
	for i := 0; i < 100; i++ {               //run 100 tests to be sure if it is prime or not
		if Fermat(argString) == false {
			myMap[false]++ //increment false
		} else {
			myMap[true]++ //increment true
		}
	}
	if myMap[false] > myMap[true] { //if false count is higher than true count
		fmt.Println("The number is not prime")
	} else {
		fmt.Println("The number is prime")
	}
}

//Fermat tests primality using Fermat's Little Theorem
func Fermat(num int) bool {
	rand.Seed(time.Now().UTC().UnixNano())            //seed random number to time
	randomNum := rand.Intn(num-1) + 1                 //pick a random number between 1 and the number
	x := math.Pow(float64(randomNum), float64(num-1)) //calculate randomNumber^(num-1)
	if int(x)%num != 1 {                              //if randomNumber^(num-1) is not equal to 1
		return false //it is not a prime number
	}
	return true //if it is equal to 1 it is prime

}
