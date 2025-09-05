// p0875_koko_eating_bananas p0875 koko eating bananas
// Link: https://leetcode.com/problems/p0875_koko_eating_bananas/
// Pattern: TBD
// Invariants: TBD

package p0875_koko_eating_bananas

// Usolve solves the problem
// Add parameters and return types according to the problem requirements
func minEatingSpeed(piles []int, h int) int {
	l, r := 1, -1
	for _, pile := range piles {
		r = max(r, pile)
	}
	for l < r {
		mid := l + (r-l)/2
		c := valid(piles, h, mid)
		if c == h {
			r = mid
		} else if c < h {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

func valid(piles []int, h, k int) int {
	t := 0
	for _, pile := range piles {
		t += pile / k
		if pile%k > 0 {
			t++
		}
	}
	return t
}
