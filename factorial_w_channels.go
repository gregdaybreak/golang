package main

import (
	"fmt"
)

func main() {
	c := factorial(4)   //pass 4 to the factorial function and get back the channel
	for n := range c {  //range over the channel to receive the factorial
		fmt.Println(n)   //print the result
	}
}

//factorial figures out the factorial of a given integer and returns a channel of type int
func factorial(num int) chan int {
	out := make(chan int)   //make a new channel
	go func() {    //start the goroutine to calculate the factorial
		total := 1
		for i := num; i > 0; i-- {
			total *= i
		}
		out <- total   //send total to the channel named out
		close(out) //close the channel
	}()
	return out  //return the channel
}
