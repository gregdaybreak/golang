package main

import (
	"fmt"
	"sort"
)

//this creates a 6x6 2D Array and calculates the sum's of an hourglass in the array.
//it then prints out the largest sum of the hourglass.
func main() {
	var row, col int = 6, 6 //create our 6x6 slice
	arr := make([][]int, row)
	for i := range arr {
		arr[i] = make([]int, col)
	}
	for i := 0; i < 6; i++ { //populate our 6x6 2d slice from stdin
		for j := 0; j < 6; j++ {
			var input int
			fmt.Scan(&input)
			arr[i][j] = input
		}
	}
	sum := make([]int, 16) //make a slice for the sums
	var h int
	for i := 0; i < 4; i++ { //hourglasses are 3x3
		for j := 0; j < 4; j++ {
			sum[h] = arr[i][j] + arr[i][j+1] + arr[i][j+2] + arr[i+1][j+1] + arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2] //add up the cells that make up the hourglasses
			h++
		}
	}
	sort.Ints(sum)            //sort the slice
	fmt.Printf("%d", sum[15]) //print out the biggest number
}
