// 1 Two Sum
// Link: https://leetcode.com/problems/two-sum/
// Pattern: hash map (one-pass)
// Invariants: when scanning i, all j < i are indexed; check if target-nums[i] exists.

package p0001_two_sum

// TwoSum finds two numbers in nums that add up to target and returns their indices.
// Time complexity: O(n)
// Space complexity: O(n)
func TwoSum(nums []int, target int) []int {
	// Create a map to store value -> index mappings
	numMap := make(map[int]int)

	// Iterate through the array
	for i, num := range nums {
		// Calculate the complement needed to reach target
		complement := target - num

		// Check if the complement exists in our map
		if j, exists := numMap[complement]; exists {
			// Return the indices if found
			return []int{j, i}
		}

		// Store the current number and its index
		numMap[num] = i
	}

	// Return empty slice if no solution found (shouldn't happen per problem constraints)
	return []int{}
}
