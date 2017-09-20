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
	arg, _ := strconv.ParseInt(os.Args[1], 10, 64) //convert argument to int
	myMap := map[bool]int{true: 0, false: 0}       //new map to hold bool count
	for i := 0; i < 100; i++ {                     //run 100 tests to be sure if it is prime or not
		if Fermat(arg) == false { //run the Fermat function and get result
			myMap[false]++ //increment false
		} else {
			myMap[true]++ //increment true
		}
	}
	fmt.Println(myMap)
	if myMap[true] > 30 { //if true count is greater than 30 we are pretty confident it is prime
		fmt.Printf("The number %v is prime\n", arg)
	} else {
		fmt.Printf("The number %v is not prime\n", arg)
	}
}

//Fermat tests primality using Fermat's Little Theorem
func Fermat(num int64) bool {
	rand.Seed(time.Now().UTC().UnixNano())            //seed random number to time
	randomNum := rand.Int63n(num-1) + 1               //pick a random number between 1 and the number
	x := math.Pow(float64(randomNum), float64(num-1)) //calculate randomNumber^(num-1)
	if int64(x)%num != 1 {                            //if randomNumber^(num-1) is not equal to 1
		return false //it is not a prime number
	}
	return true //if it is equal to 1 it is prime

}
