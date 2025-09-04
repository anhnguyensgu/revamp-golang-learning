// p739_daily_temperatures p739 daily temperatures
// Link: https://leetcode.com/problems/p739_daily_temperatures/
// Pattern: TBD
// Invariants: TBD

package p0739_daily_temperatures

// Usolve solves the problem
// Add parameters and return types according to the problem requirements
func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))

	st := []int{}
	for i, t := range temperatures {
		for len(st) > 0 && t > temperatures[st[len(st)-1]] {
			j := st[len(st)-1]
			st = st[:len(st)-1]
			ans[j] = i - j
		}
		st = append(st, i)

	}
	return ans
}
