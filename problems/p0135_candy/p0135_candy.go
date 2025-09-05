// p0135_candy p0135 candy
// Link: https://leetcode.com/problems/p0135_candy/
// Pattern: TBD
// Invariants: TBD

package p0135_candy

// Usolve solves the problem
// Add parameters and return types according to the problem requirements
func candy(ratings []int) int {
	ans := 0
	candies := make([]int, len(ratings))
	for i := range candies {
		candies[i] = 1
	}
	for i := 1; i < len(ratings); i++ {
		if ratings[i-1] < ratings[i] {
			candies[i] = candies[i-1] + 1
		}
	}

	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = max(candies[i], candies[i+1]+1)
		}
	}

	for _, count := range candies {
		ans += count
	}
	return ans
}
