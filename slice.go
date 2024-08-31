// Package goList 提供了对泛型切片的常见操作。
package goList

type slice[T any] struct {
	slice  []T
	Length int
}

// New 创建并返回一个新的 Slice 实例。
// 如果传入一个正整数作为唯一参数，则返回一个具有预设容量的空切片。
// 如果传入多个元素，则这些元素将被添加到新的切片中。
func New[T any](elements ...T) *slice[T] {
	if len(elements) == 1 {
		if value, ok := any(elements[0]).(int); ok && value > 0 {
			return &slice[T]{
				slice:  make([]T, 0, value),
				Length: value,
			}
		}
	}
	result := append([]T{}, elements...)

	return &slice[T]{
		slice:  result,
		Length: len(result),
	}
}
