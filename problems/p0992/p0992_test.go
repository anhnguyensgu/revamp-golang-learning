package p0992

import (
	"testing"
)

func TestSubarraysWithKDistinct(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 1, 2, 3},
			k:        2,
			expected: 7,
		},
		{
			name:     "Example 2",
			nums:     []int{1, 2, 1, 3, 4},
			k:        3,
			expected: 3,
		},
		{
			name:     "Single element k=1",
			nums:     []int{1},
			k:        1,
			expected: 1,
		},
		{
			name:     "All same elements",
			nums:     []int{1, 1, 1, 1},
			k:        1,
			expected: 10,
		},
		{
			name:     "k larger than distinct",
			nums:     []int{1, 2, 3},
			k:        4,
			expected: 0,
		},
		{
			name:     "k equals length",
			nums:     []int{1, 2, 3, 4},
			k:        4,
			expected: 1,
		},
		{
			name: "k equals length",
			nums: []int{1, 2, 1, 2, 1},
			// nums:      []int{0, 0, 1, 2, 3},
			k:        2,
			expected: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := subarraysWithKDistinct(tt.nums, tt.k)
			if result != tt.expected {
				t.Errorf("subarraysWithKDistinct(%v, %d) = %d, expected %d", tt.nums, tt.k, result, tt.expected)
			}
		})
	}
}
