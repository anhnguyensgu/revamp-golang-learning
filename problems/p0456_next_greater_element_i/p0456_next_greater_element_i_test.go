package p0123_next_greater_element_i

import (
	"reflect"
	"testing"
)

func TestNextGreaterElement(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  []int
	}{
		// Basic cases
		{"empty_nums1", []int{}, []int{1, 2, 3}, []int{}},
		{"single_element_found", []int{1}, []int{1, 2}, []int{2}},
		{"single_element_not_found", []int{2}, []int{1, 2}, []int{-1}},

		// Provided example
		{"provided_example", []int{4, 1, 2}, []int{1, 3, 4, 2}, []int{-1, 3, -1}},
		{"provided_example", []int{2, 4}, []int{1, 2, 3, 4}, []int{3, -1}},
		//Input: nums1 = [2,4], nums2 = [1,2,3,4]
		//Output: [3,-1]

		// Ordering patterns
		{"increasing_sequence", []int{1, 2, 3}, []int{1, 2, 3, 4}, []int{2, 3, 4}},
		{"decreasing_sequence", []int{4, 3, 2, 1}, []int{1, 2, 3, 4}, []int{-1, 4, 3, 2}},
		{"mixed_pattern", []int{2, 4, 1}, []int{1, 2, 3, 4}, []int{3, -1, 2}},

		// All elements scenarios
		{"all_have_next_greater", []int{1, 2, 3}, []int{1, 2, 3, 5}, []int{2, 3, 5}},
		{"none_have_next_greater", []int{5, 4, 3}, []int{1, 2, 3, 4, 5}, []int{-1, 5, 4}},

		// Duplicates handling
		{"duplicates_with_greater", []int{2, 2}, []int{2, 2, 3}, []int{3, 3}},
		{"all_same_values", []int{5, 5}, []int{5, 5, 5}, []int{-1, -1}},

		// Edge values within constraints
		{"min_values", []int{0, 1}, []int{0, 1, 2}, []int{1, 2}},
		{"large_values", []int{9999, 10000}, []int{9999, 10000, 10001}, []int{10000, 10001}},

		// Complex patterns
		{"mountain_pattern", []int{1, 3, 2}, []int{1, 2, 3, 4}, []int{2, 4, 3}},
		{"valley_pattern", []int{3, 1, 4}, []int{1, 3, 4, 5}, []int{4, 3, 5}},
		{"random_order", []int{2, 7, 3}, []int{2, 7, 3, 5, 4, 6, 8}, []int{7, 8, 5}},

		// Maximum length scenario (partial)
		{"larger_nums1", []int{1, 5, 10, 15, 20}, []int{1, 2, 5, 8, 10, 12, 15, 18, 20, 25}, []int{2, 8, 12, 18, 25}},

		// Sparse distribution
		{"sparse_elements", []int{100, 500, 900}, []int{50, 100, 200, 500, 600, 900, 1000}, []int{200, 600, 1000}},

		// No next greater at end
		{"last_elements_largest", []int{900, 1000}, []int{100, 500, 900, 1000}, []int{1000, -1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := nextGreaterElement(tt.nums1, tt.nums2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElement(%v, %v) = %v, want %v", tt.nums1, tt.nums2, got, tt.want)
			}
		})
	}
}
