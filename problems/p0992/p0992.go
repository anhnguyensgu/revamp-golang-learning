// p0992 p0992
// Link: https://leetcode.com/problems/p0992/
// Pattern: TBD
// Invariants: TBD

//992. Subarrays with K Different Integers
//Hard
//Topics
//premium lock icon
//Companies
//Hint
//Given an integer array nums and an integer k, return the number of good subarrays of nums.
//
//A good array is an array where the number of different integers in that array is exactly k.
//
//For example, [1,2,3,1,2] has 3 different integers: 1, 2, and 3.
//A subarray is a contiguous part of an array.
//
//
//
//Example 1:
//
//Input: nums = [1,2,1,2,3], k = 2
//Output: 7
//Explanation: Subarrays formed with exactly 2 different integers: [1,2], [2,1], [1,2], [2,3], [1,2,1], [2,1,2], [1,2,1,2]
//Example 2:
//
//Input: nums = [1,2,1,3,4], k = 3
//Output: 3
//Explanation: Subarrays formed with exactly 3 different integers: [1,2,1,3], [2,1,3], [1,3,4].
//
//
//Constraints:
//
//1 <= nums.length <= 2 * 104
//1 <= nums[i], k <= nums.length

package p0992

// Usolve solves the problem in a single pass using two sliding windows that track
// subarrays with at most k and at most (k-1) distinct values respectively.
// The difference between their left bounds yields the count of subarrays ending
// at the current right index that have exactly k distinct values.
func subarraysWithKDistinct(nums []int, k int) int {
	ans := 0
	left := 0
	distinctCount := make([]int, len(nums)+1)
	curCount := 0
	for _, value := range nums {
		distinctCount[value]++
		if distinctCount[value] == 1 {
			k--
		}

		if k < 0 {
			distinctCount[nums[left]]--
			left++
			k++
			curCount = 0
		}

		if k == 0 {
			for distinctCount[nums[left]] > 1 {
				distinctCount[nums[left]]--
				left++
				curCount++
			}
			ans += (curCount + 1)
		}
	}
	return ans
}
