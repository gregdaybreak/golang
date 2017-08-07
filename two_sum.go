package main

func main() {
	mySlice := []int{2, 7, 11, 15}
	myLength := 17
	twoSum(mySlice, myLength)
}

//TwoSum takes a slice of integers and a target sum and returns a slice of int
//that says which index values of the slice made up that sum
func twoSum(nums []int, target int) []int {
	length := len(nums) //assign the number of elements in the slice to variable length

	resMap := make(map[int]int, length) //create a new map with a key and value of int and a length and capacity of the length of the slice

	for i := 0; i < length; i++ { //while i is less that the length of the slice
		c := target - nums[i]       // c equals the target minus the value in the slice of int
		if _, ok := resMap[c]; ok { //if the key for the target minus the value exists
			return []int{resMap[c], i} //return a slice of int containing the index where the value lives and the index
		} else { //otherwise
			resMap[nums[i]] = i //store into resMap with the value in the slice as the key and the index location as the value
		}
	}

	return nil //dont return the value in this case
}
