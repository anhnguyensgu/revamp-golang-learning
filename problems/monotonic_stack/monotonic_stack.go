// monotonic_stack monotonic stack
// Link: https://leetcode.com/problems/monotonic_stack/
// Pattern: TBD
// Invariants: TBD

package monotonic_stack

// NextGreaterElement solves the problem
// Add parameters and return types according to the problem requirements
func NextGreaterElement(arr []int) []int {
	ans := make([]int, len(arr))
	for i := range ans {
		ans[i] = -1
	}

	st := []int{}

	for i, cur := range arr {
		for len(st) > 0 && cur > arr[st[len(st)-1]] {
			// pop index
			j := st[len(st)-1]
			st = st[:len(st)-1]
			ans[j] = arr[i]
		}
		st = append(st, i)
	}

	return ans
}
