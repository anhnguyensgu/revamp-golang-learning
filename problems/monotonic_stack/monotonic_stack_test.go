package monotonic_stack

import (
	"reflect"
	"testing"
)

func TestNextGreaterElement(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{"empty", []int{}, []int{}},
		{"single", []int{1}, []int{-1}},
		{"increasing", []int{1, 2, 3, 4}, []int{2, 3, 4, -1}},
		{"decreasing", []int{4, 3, 2, 1}, []int{-1, -1, -1, -1}},
		{"mixed", []int{4, 3, 1, 2}, []int{-1, -1, 2, -1}},
		{"duplicates_then_greater", []int{2, 2, 3}, []int{3, 3, -1}},
		{"all_equal", []int{1, 1, 1}, []int{-1, -1, -1}},
		{"negatives", []int{-2, -1, -3}, []int{-1, -1, -1}},
		{"random", []int{2, 7, 3, 5, 4, 6, 8}, []int{7, 8, 5, 6, 6, 8, -1}},
	}

	for _, tt := range tests {
		got := NextGreaterElement(tt.input)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("%s: NextGreaterElement(%v) = %v, want %v", tt.name, tt.input, got, tt.want)
		}
	}
}
