// p0123_next_greater_element_i p0123 next greater element i
// Link: https://leetcode.com/problems/p0123_next_greater_element_i/
// Pattern: TBD
// Invariants: TBD

package p0123_next_greater_element_i

// nextGreaterElement solves the problem
// Add parameters and return types according to the problem requirements
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var st []int
	next := map[int]int{}
	for i := 0; i < len(nums2); i++ {
		for len(st) > 0 && nums2[i] > nums2[st[len(st)-1]] {
			j := st[len(st)-1]
			st = st[:len(st)-1]
			next[nums2[j]] = nums2[i]
		}
		st = append(st, i)
	}
	for i, cur := range nums1 {
		if value, ok := next[cur]; ok {
			nums1[i] = value
		} else {
			nums1[i] = -1
		}
	}

	return nums1
}
