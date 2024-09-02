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

func TestSliceFlat(t *testing.T) {
	tests := []struct {
		name      string
		input     []any
		deep      []int
		expected  []any
		expectErr bool
	}{
		{
			name:      "Flat array with default depth (1)",
			input:     []any{1, []any{2, 3}, 4},
			deep:      []int{},
			expected:  []any{1, 2, 3, 4},
			expectErr: false,
		},
		{
			name:      "Flat array with depth 2",
			input:     []any{1, []any{2, []any{3, 4}}, 5},
			deep:      []int{},
			expected:  []any{1, 2, []any{3, 4}, 5},
			expectErr: false,
		},
		{
			name:      "Flat array with depth 2",
			input:     []any{1, []any{2, []any{3, 4}}, 5},
			deep:      []int{2},
			expected:  []any{1, 2, 3, 4, 5},
			expectErr: false,
		},
		{
			name:      "Flat array with depth 0 (no flattening)",
			input:     []any{1, []any{2, []any{3, 4}}, 5},
			deep:      []int{0},
			expected:  []any{1, []any{2, []any{3, 4}}, 5},
			expectErr: false,
		},
		{
			name:      "Invalid deep parameter (more than 1)",
			input:     []any{1, 2, 3},
			deep:      []int{1, 2},
			expected:  nil,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var args []any
			for _, v := range tt.input {
				args = append(args, v)
			}

			s := New(args...)
			result, err := s.Flat(tt.deep...)

			if (err != nil) != tt.expectErr {
				t.Errorf("Flat() error = %v, expectErr %v", err, tt.expectErr)
				return
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Flat() = %v, expected %v", result, tt.expected)
			}
		})
	}
}
