package main

import (
	"fmt"
	"math"
	"strconv"
)

//converts an integer from standin to binary and then calculates how many consecutive 1's are in the binary representation and prints it.
func main() {

	var num int
	fmt.Scan(&num)                             //scan in integer
	binary := strconv.FormatInt(int64(num), 2) //convert num to int64 to convert to binary
	var maxConsecutiveOneNum int
	var counter int
	for i := 0; i < len(binary); i++ { //loop over the number in binary for as long as the binary number is
		if string([]rune(binary)[i]) == "1" { //if the binary value is a 1
			counter++                                                                             //increment
			maxConsecutiveOneNum = int(math.Max(float64(maxConsecutiveOneNum), float64(counter))) //returns the larger of the two nums
		} else {
			counter = 0 //set the counter back to 0 if the 1's are not consecutive
		}
	}
	fmt.Printf("%d", maxConsecutiveOneNum)
}
