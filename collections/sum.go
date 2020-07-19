package main

// Sum returns the sum by array
func Sum(nums []int) int {
	var result int
	for _, num := range nums {
		result += num
	}
	return result
}

// SumAll returns slice of sum of the slices passed as parameters
func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

// SumAllTails returns slice of sum of the slices passed as parameters except for the first element
func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(numbers[1:]))
		}
	}
	return sums
}
