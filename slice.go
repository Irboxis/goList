// Package goList 提供了对泛型切片的常见操作。
package goList

import "errors"

type slice[T any] struct {
	Slice  []T
	Length int
	zero   T
}

// New 创建并返回一个新的 slice 实例。
// 如果传入一个正整数作为唯一参数，则返回一个具有预设容量的空切片。
// 如果传入多个元素，则这些元素将被添加到新的切片中。
func New[T any](elements ...T) *slice[T] {
	if len(elements) == 1 {
		if value, ok := any(elements[0]).(int); ok && value > 0 {
			return &slice[T]{
				Slice:  make([]T, 0, value),
				Length: value,
			}
		}
	}
	result := append([]T{}, elements...)

	return &slice[T]{
		Slice:  result,
		Length: len(result),
	}
}

func (s *slice[T]) isNullS() bool {
	if s.Length == 0 {
		return false
	}

	return true
}

func indexCheck[T any](s *slice[T], index int) (int, error) {
	if s == nil {
		return -1, errors.New("slice is nil")
	}

	if index < 0 {
		index += s.Length
	}

	if index > s.Length || index < 0 {
		return -1, errors.New("unexpected argument, index should be within slice bounds")
	}
	return index, nil
}
