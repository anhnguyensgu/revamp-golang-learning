// p084_largest_rectangle_in_histogram p084 largest rectangle in histogram
// Link: https://leetcode.com/problems/p084_largest_rectangle_in_histogram/
// Pattern: TBD
// Invariants: TBD

package p084_largest_rectangle_in_histogram

func largestRectangleArea(heights []int) int {
	ans := 0
	left := make([]int, len(heights))
	for i := 0; i < len(heights); i++ {
		left[i] = -1
	}
	var stack []int
	//closest smaller
	for i, a := range heights {
		for len(stack) > 0 && a <= heights[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	right := make([]int, len(heights))
	for i := 0; i < len(heights); i++ {
		right[i] = len(heights)
	}

	stack = []int{}

	for i := len(heights) - 1; i >= 0; i-- {
		for len(stack) > 0 && heights[i] <= heights[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			right[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	for i := 0; i < len(heights); i++ {
		area := (right[i] - left[i] - 1) * heights[i]
		ans = max(ans, area)
	}
	return ans
}
