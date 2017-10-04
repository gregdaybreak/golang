package main

import "fmt"

var a = []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

func main() {

	insertionSort(a)
	fmt.Println(a)
	b := []int{5, 4, 3, 2, 1}
	insertionSort(b)
	fmt.Println(b)
}

//insertionSort sorts a slice using an insertion sort algorithm.
func insertionSort(arr []int) {

	for j := 1; j < len(arr); j++ { //while j is less than the length of the slice
		key := arr[j]                //assign the value at the index of the array to key (this is what we are actively sorting.)
		i := j - 1                   //i is the index one less than j
		for i >= 0 && arr[i] > key { //while i is greater than or equal to 0 and the value
			//of the previous item is larger than the value in front of it (it needs to be sorted)
			arr[i+1] = arr[i] //the larger value is assigned to the current location
			i--               //allows us to break the loop since it goes below 0 when the sort is done
		}
		arr[i+1] = key //the key gets now gets assigned to the previous spot.
	}
}
