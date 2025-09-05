package p0135_candy

import (
	"testing"
)

func TestCandy(t *testing.T) {
	testCases := []struct {
		name     string
		ratings  []int
		expected int
	}{
		{
			name:     "Example 1: [1,0,2]",
			ratings:  []int{1, 0, 2},
			expected: 5,
		},
		{
			name:     "Example 2: [1,2,2]",
			ratings:  []int{1, 2, 2},
			expected: 4,
		},
		{
			name:     "Single child",
			ratings:  []int{5},
			expected: 1,
		},
		{
			name:     "All same ratings",
			ratings:  []int{3, 3, 3, 3},
			expected: 4,
		},
		{
			name:     "Increasing sequence",
			ratings:  []int{1, 2, 3, 4, 5},
			expected: 15, // 1+2+3+4+5
		},
		{
			name:     "Decreasing sequence",
			ratings:  []int{5, 4, 3, 2, 1},
			expected: 15, // 5+4+3+2+1
		},
		{
			name:     "Mountain shape",
			ratings:  []int{1, 2, 3, 2, 1},
			expected: 9, // 1+2+3+2+1
		},
		{
			name:     "Valley shape",
			ratings:  []int{3, 2, 1, 2, 3},
			expected: 11, // 3+2+1+2+3
		},
		{
			name:     "Complex case",
			ratings:  []int{1, 3, 2, 2, 1},
			expected: 7, // 1+2+1+2+1
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := candy(tc.ratings)
			if result != tc.expected {
				t.Errorf("Expected %d, got %d for ratings %v", tc.expected, result, tc.ratings)
			}
		})
	}
}
