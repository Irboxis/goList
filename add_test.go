package goList

import (
	"reflect"
	"testing"
)

func TestSlice_Add(t *testing.T) {
	s := &list[int]{
		Slice:  []int{1, 2, 4},
		Length: 3,
	}

	err := s.Add(2, 3)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	expected := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(s.Slice, expected) {
		t.Errorf("expected slice %v, got %v", expected, s.Slice)
	}
	if s.Length != 4 {
		t.Errorf("expected length 4, got %d", s.Length)
	}

	// Test with negative index
	err = s.Add(-2, 5)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	expected = []int{1, 2, 5, 3, 4}
	if !reflect.DeepEqual(s.Slice, expected) {
		t.Errorf("expected slice %v, got %v", expected, s.Slice)
	}
	if s.Length != 5 {
		t.Errorf("expected length 5, got %d", s.Length)
	}

	// Test with out-of-bounds index
	err = s.Add(10, 6)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestSlice_AddAll(t *testing.T) {
	s := &list[int]{
		Slice:  []int{1, 2, 5},
		Length: 3,
	}

	err := s.AddAll(2, 3, 4)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	expected := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(s.Slice, expected) {
		t.Errorf("expected slice %v, got %v", expected, s.Slice)
	}
	if s.Length != 5 {
		t.Errorf("expected length 5, got %d", s.Length)
	}

	// Test with negative index
	err = s.AddAll(-1, 6, 7)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	expected = []int{1, 2, 3, 4, 6, 7, 5}
	if !reflect.DeepEqual(s.Slice, expected) {
		t.Errorf("expected slice %v, got %v", expected, s.Slice)
	}
	if s.Length != 7 {
		t.Errorf("expected length 7, got %d", s.Length)
	}

	// Test with out-of-bounds index
	err = s.AddAll(10, 8, 9)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}
