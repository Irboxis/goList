package goList

import (
	"testing"
)

func TestIndexOf(t *testing.T) {
	s := New(1, 2, 3, 4, 5)

	index, err := s.IndexOf(3)
	if err != nil || index != 2 {
		t.Errorf("Expected index 2, got %d, error: %v", index, err)
	}

	index, err = s.IndexOf(6)
	if err == nil || index != -1 {
		t.Errorf("Expected error for element not found, got index %d", index)
	}

	emptySlice := New[int]()
	index, err = emptySlice.IndexOf(1)
	if err == nil || index != -1 {
		t.Errorf("Expected error for empty slice, got index %d", index)
	}
}

func TestValueOf(t *testing.T) {
	s := New(1, 2, 3, 4, 5)

	val, err := s.ValueOf(2)
	if err != nil || val != 3 {
		t.Errorf("Expected value 3, got %d, error: %v", val, err)
	}

	val, err = s.ValueOf(-1)
	if err != nil || val != 5 {
		t.Errorf("Expected value 5 for index -1, got %d, error: %v", val, err)
	}

	val, err = s.ValueOf(10)
	if err != nil || val != 5 {
		t.Errorf("Expected value 5 for index out of bounds, got %d, error: %v", val, err)
	}

	emptySlice := New[int]()
	val, err = emptySlice.ValueOf(0)
	if err == nil || val != 0 {
		t.Errorf("Expected error for empty slice, got value %d", val)
	}
}

func TestFind(t *testing.T) {
	s := New(1, 2, 3, 4, 5)

	val, err := s.Find(func(elem int, index int, slice []int) bool {
		return elem > 3
	})
	if err != nil || val != 4 {
		t.Errorf("Expected value 4, got %d, error: %v", val, err)
	}

	val, err = s.Find(func(elem int, index int, slice []int) bool {
		return elem > 6
	})
	if err == nil || val != 0 {
		t.Errorf("Expected error for element not found, got value %d", val)
	}

	emptySlice := New[int]()
	val, err = emptySlice.Find(func(elem int, index int, slice []int) bool {
		return elem > 0
	})
	if err == nil || val != 0 {
		t.Errorf("Expected error for empty slice, got value %d", val)
	}
}

func TestIncludes(t *testing.T) {
	s := New(1, 2, 3, 4, 5)

	exists, err := s.Includes(3)
	if err != nil || !exists {
		t.Errorf("Expected true for existing element, got %v, error: %v", exists, err)
	}

	exists, err = s.Includes(6)
	if err != nil || exists {
		t.Errorf("Expected false for non-existing element, got %v, error: %v", exists, err)
	}

	exists, err = s.Includes(3, 2)
	if err != nil || !exists {
		t.Errorf("Expected true for existing element starting from index 2, got %v, error: %v", exists, err)
	}

	exists, err = s.Includes(3, 4)
	if err != nil || exists {
		t.Errorf("Expected false for non-existing element starting from index 4, got %v, error: %v", exists, err)
	}

	emptySlice := New[int]()
	exists, err = emptySlice.Includes(1)
	if err == nil || exists {
		t.Errorf("Expected error for empty slice, got %v", exists)
	}

	exists, err = s.Includes(3, 1, 2)
	if err == nil || exists {
		t.Errorf("Expected error for multiple start values, got %v", exists)
	}
}
