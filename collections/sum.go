package main

// Sum returns the sum by array
func Sum(nums []int) int {
	var result int
	for _, num := range nums {
		result += num
	}
	return result
}
