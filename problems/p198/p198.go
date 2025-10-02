// p198 p198
// Link: https://leetcode.com/problems/p198/
// Pattern: TBD
// Invariants: TBD
package p198

// You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected and it will automatically contact the police if two adjacent houses were broken into on the same night.
//
// Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.
//
//
//
// Example 1:
//
// Input: nums = [1,2,3,1]
// Output: 4
// Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
// Total amount you can rob = 1 + 3 = 4.
// Example 2:
//
// Input: nums = [2,7,9,3,1]
// Output: 12
// Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
// Total amount you caConstraints:

// 1 <= nums.length <= 100
// 0 <= nums[i] <= 400n rob = 2 + 9 + 1 = 12.

// Usolve solves the problem
// Add parameters and return types according to the problem requirements
// [2,1,1,2]
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	first, second := nums[0], nums[1]
	for i := 2; i < len(nums); i++ {
		tmp := nums[i] + first
		first = max(second, first)
		second = tmp
	}

	return max(first, second)
}

func robWithCircle(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	if len(nums) == 3 {
		return max(nums[2], max(nums[0], nums[1]))
	}

	first, second := nums[0], nums[1]

	for i := 2; i < len(nums)-1; i++ {
		tmp := nums[i] + first
		first = max(second, first)
		second = tmp
	}

	ans := max(first, second)

	first, second = nums[1], nums[2]
	for i := 3; i < len(nums); i++ {
		tmp := nums[i] + first
		first = max(second, first)
		second = tmp
	}

	return max(ans, max(first, second))
}

// [1,1,3,6,7,10,7,1,8,5,9,1,4,4,3]

// You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed. All houses at this place are arranged in a circle. That means the first house is the neighbor of the last one. Meanwhile, adjacent houses have a security system connected, and it will automatically contact the police if two adjacent houses were broken into on the same night.
//
// Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.
//
//
//
// Example 1:
//
// Input: nums = [2,3,2]
// Output: 3
// Explanation: You cannot rob house 1 (money = 2) and then rob house 3 (money = 2), because they are adjacent houses.
// Example 2:
//
// Input: nums = [1,2,3,1]
// Output: 4
// Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
// Total amount you can rob = 1 + 3 = 4.
// Example 3:
//
// Input: nums = [1,2,3]
// Output: 3
//
//
// Constraints:
//
// 1 <= nums.length <= 100
// 0 <= nums[i] <= 1000

//how to track the first and last element
