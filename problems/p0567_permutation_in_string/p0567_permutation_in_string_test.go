package p0567_permutation_in_string

import (
	"testing"
)

func TestUsolve(t *testing.T) {
}

func TestCheckInclusion(t *testing.T) {
	tests := []struct {
		name     string
		s1       string
		s2       string
		expected bool
	}{
		{
			name:     "Example 1: permutation exists",
			s1:       "ab",
			s2:       "eidbaooo",
			expected: true,
		},
		{
			name:     "Example 2: permutation does not exist",
			s1:       "ab",
			s2:       "eidboaoo",
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := checkInclusion(test.s1, test.s2)
			if result != test.expected {
				t.Errorf("checkInclusion(%q, %q) = %v, expected %v", test.s1, test.s2, result, test.expected)
			}
		})
	}
}
