// p0567_permutation_in_string p0567 permutation in string
// Link: https://leetcode.com/problems/p0567_permutation_in_string/
// Pattern: TBD
// Invariants: TBD

package p0567_permutation_in_string

// Usolve solves the problem
// Add parameters and return types according to the problem requirements
func checkInclusion(s1 string, s2 string) bool {
	charFre := [26]int{}
	for _, char := range s1 {
		charFre[char-'a']++
	}

	charFre2 := [26]int{}

	j := 0
	for i, char := range s2 {
		charFre2[char-'a']++
		l := i - j + 1
		if l > len(s1) {
			charFre2[s2[j]-'a']--
			j++
		}

		if i-j+1 == len(s1) {
			isMatch := true
			for idx := range charFre {
				if charFre2[idx] != charFre[idx] {
					isMatch = false
					break
				}
			}
			if isMatch {
				return true
			}
		}
	}
	return false
}
