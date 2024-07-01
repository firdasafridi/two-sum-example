// main.go
package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	// Create a map to store the index of the elements
	numMap := make(map[int]int)

	// Iterate through the array
	for i, num := range nums {
		// Calculate the complement
		complement := target - num

		// Check if the complement exists in the map
		if index, ok := numMap[complement]; ok {
			// If found, return the indices
			return []int{index, i}
		}

		// Store the index of the current number
		numMap[num] = i
	}

	// Return an empty slice if no solution is found (though the problem guarantees one solution)
	return []int{}
}

func main() {
	// Sample input
	nums := []int{2, 7, 11, 15}
	target := 9

	// Get the result
	result := twoSum(nums, target)

	// Print the result
	fmt.Println(result)
}
