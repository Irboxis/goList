package goList

import (
	"reflect"
	"testing"
)

func TestDelete(t *testing.T) {
	// 测试删除单个元素
	s := &list[int]{Slice: []int{1, 2, 3, 4, 5}, Length: 5}
	deleted, err := s.Delete(1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	expected := []int{2}
	if !reflect.DeepEqual(deleted, expected) {
		t.Errorf("expected %v, got %v", expected, deleted)
	}
	expectedSlice := []int{1, 3, 4, 5}
	if !reflect.DeepEqual(s.Slice, expectedSlice) {
		t.Errorf("expected slice %v, got %v", expectedSlice, s.Slice)
	}

	// 测试删除多个元素
	s = &list[int]{Slice: []int{1, 2, 3, 4, 5}, Length: 5}
	deleted, err = s.Delete(1, 2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	expected = []int{2, 3}
	if !reflect.DeepEqual(deleted, expected) {
		t.Errorf("expected %v, got %v", expected, deleted)
	}
	expectedSlice = []int{1, 4, 5}
	if !reflect.DeepEqual(s.Slice, expectedSlice) {
		t.Errorf("expected slice %v, got %v", expectedSlice, s.Slice)
	}
}

func TestPop(t *testing.T) {
	s := &list[int]{Slice: []int{1, 2, 3, 4, 5}, Length: 5}
	popped, _ := s.Pop()
	expected := 5
	if popped != expected {
		t.Errorf("expected %v, got %v", expected, popped)
	}
	expectedSlice := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(s.Slice, expectedSlice) {
		t.Errorf("expected slice %v, got %v", expectedSlice, s.Slice)
	}
}

func TestUnShift(t *testing.T) {
	s := &list[int]{Slice: []int{1, 2, 3, 4, 5}, Length: 5}
	unsifted, _ := s.UnShift()
	expected := 1
	if unsifted != expected {
		t.Errorf("expected %v, got %v", expected, unsifted)
	}
	expectedSlice := []int{2, 3, 4, 5}
	if !reflect.DeepEqual(s.Slice, expectedSlice) {
		t.Errorf("expected slice %v, got %v", expectedSlice, s.Slice)
	}
}
