package goList

import (
	"testing"
)

func TestNewIntSlice(t *testing.T) {
	// 测试创建包含整数的 slice
	intSlice := New(1, 2, 3, 4, 5)

	if intSlice.Length != 5 {
		t.Errorf("Expected length 5, got %d", intSlice.Length)
	}

	expected := []int{1, 2, 3, 4, 5}
	for i, v := range intSlice.slice {
		if v != expected[i] {
			t.Errorf("Expected %d at index %d, got %d", expected[i], i, v)
		}
	}
}

func TestNewEmptySlice(t *testing.T) {
	LSlice := New(3)

	if LSlice.Length != 3 {
		t.Errorf("Expected length 0, got %d", LSlice.Length)
	}
}
