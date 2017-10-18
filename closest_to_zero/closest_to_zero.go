package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

//solution to code in game's Temperatures
func main() {
	var n int
	fmt.Scanf("%d", &n) //scan in the first integer (number of temps to analyze)

	reader := bufio.NewReader(os.Stdin)
	temps, _ := reader.ReadString('\n') //read in the strings until you see a newline
	tokens := strings.Split(temps, " ") //split on whitespace and insert into a slice

	if len(tokens) == 0 { //if the length of the slice is "0"
		fmt.Println("0") //print "0"
		os.Exit(0)       //and exit
	}

	c := 5527                          //c is equal to 5527 initially (this is bigger than the temp range given)
	for i := 0; i < len(tokens); i++ { //go through the slice
		intval, _ := strconv.Atoi(strings.TrimSpace(tokens[i])) //clean all whitespace with TrimSpace, convert string to int
		if math.Abs(float64(intval)) < math.Abs(float64(c)) || (math.Abs(float64(intval)) == math.Abs(float64(c)) && intval > c) {
			//if the number in the slice is less than 5527 initially it is closer to zero OR if the absolute value of the number is equal to c and the number is bigger than c
			c = intval //the slice value is closer to zero
		}
	}

	fmt.Println(c) //print the slice value
}
