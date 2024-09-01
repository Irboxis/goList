package goList

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name     string
		slice    *slice[int] // 被测试的切片
		toMerge  [][]int     // 要合并的多个切片
		expected []int       // 期望的合并结果
	}{
		{
			name:     "Merge with multiple non-empty slices",
			slice:    &slice[int]{Slice: []int{1, 2, 3}, Length: 3},
			toMerge:  [][]int{{4, 5}, {6, 7, 8}},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name:     "Merge with empty slice",
			slice:    &slice[int]{Slice: []int{}, Length: 0},
			toMerge:  [][]int{{1, 2, 3}},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Merge with no additional slices",
			slice:    &slice[int]{Slice: []int{1, 2, 3}, Length: 3},
			toMerge:  nil,
			expected: []int{1, 2, 3},
		},
		{
			name:     "Merge with empty input slices",
			slice:    &slice[int]{Slice: []int{1, 2}, Length: 2},
			toMerge:  [][]int{{}, {}},
			expected: []int{1, 2},
		},
		{
			name:     "Merge with both empty slice and input slices",
			slice:    &slice[int]{Slice: []int{}, Length: 0},
			toMerge:  [][]int{{}, {}},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.slice.Merge(tt.toMerge...)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("unexpected result: got %v, want %v", result, tt.expected)
			}
		})
	}
}
