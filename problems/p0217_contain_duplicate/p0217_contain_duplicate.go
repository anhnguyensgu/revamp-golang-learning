// p0217_contain_duplicate p0217 contain duplicate
// Link: https://leetcode.com/problems/p0217_contain_duplicate/
// Pattern: TBD
// Invariants: TBD

package p0217_contain_duplicate

// Usolve solves the problem
// Add parameters and return types according to the problem requirements
func containsDuplicate(nums []int) bool {
	frequency := map[int]int{}
	for _, n := range nums {
		frequency[n]++
		if frequency[n] > 1 {
			return false
		}
	}
	return true
}
