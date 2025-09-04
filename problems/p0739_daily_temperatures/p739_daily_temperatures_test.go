package p0739_daily_temperatures

import (
	"reflect"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	tests := []struct {
		name         string
		temperatures []int
		want         []int
	}{
		// Basic examples
		{
			name:         "leetcode_example_1",
			temperatures: []int{73, 74, 75, 71, 69, 72, 76, 73},
			want:         []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			name:         "ascending_sequence",
			temperatures: []int{30, 40, 50, 60},
			want:         []int{1, 1, 1, 0},
		},
		{
			name:         "mixed_ascending",
			temperatures: []int{30, 60, 90},
			want:         []int{1, 1, 0},
		},
		{
			name:         "descending_sequence",
			temperatures: []int{30, 20, 10},
			want:         []int{0, 0, 0},
		},

		// Edge cases
		{
			name:         "single_temperature",
			temperatures: []int{73},
			want:         []int{0},
		},
		{
			name:         "two_temperatures_ascending",
			temperatures: []int{30, 40},
			want:         []int{1, 0},
		},
		{
			name:         "two_temperatures_descending",
			temperatures: []int{40, 30},
			want:         []int{0, 0},
		},
		{
			name:         "all_same_temperatures",
			temperatures: []int{30, 30, 30, 30},
			want:         []int{0, 0, 0, 0},
		},

		// Complex patterns
		{
			name:         "mountain_pattern",
			temperatures: []int{30, 40, 50, 40, 30, 60},
			want:         []int{1, 1, 3, 2, 1, 0},
		},
		{
			name:         "valley_pattern",
			temperatures: []int{50, 30, 20, 40, 60},
			want:         []int{4, 2, 1, 1, 0},
		},
		{
			name:         "alternating_pattern",
			temperatures: []int{70, 60, 80, 70, 85},
			want:         []int{2, 1, 2, 1, 0},
		},

		// Large values
		{
			name:         "high_temperatures",
			temperatures: []int{100, 95, 98, 99, 101},
			want:         []int{4, 2, 1, 1, 0},
		},

		// Duplicates handling
		{
			name:         "duplicates_with_warmer",
			temperatures: []int{30, 30, 35, 30},
			want:         []int{2, 1, 0, 0},
		},
		{
			name:         "multiple_duplicates",
			temperatures: []int{73, 73, 73, 74, 73},
			want:         []int{3, 2, 1, 0, 0},
		},

		// Long sequences
		{
			name:         "longer_descending",
			temperatures: []int{80, 75, 70, 65, 60, 55},
			want:         []int{0, 0, 0, 0, 0, 0},
		},
		{
			name:         "longer_with_spike",
			temperatures: []int{70, 65, 60, 55, 50, 85, 80},
			want:         []int{5, 4, 3, 2, 1, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := dailyTemperatures(tt.temperatures)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dailyTemperatures(%v) = %v, want %v", tt.temperatures, got, tt.want)
			}
		})
	}
}
