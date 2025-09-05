// p069_sqrt p069 sqrt
// Link: https://leetcode.com/problems/p069_sqrt/
// Pattern: TBD
// Invariants: TBD

package p069_sqrt

// Usolve solves the problem
// Add parameters and return types according to the problem requirements
func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	l, r := 0, x/2 // reduce the range
	for l <= r {
		mid := l + (r-l)/2
		cur := mid * mid
		if cur == x {
			return mid
		}

		if cur > x {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	//round down
	return l - 1
}
