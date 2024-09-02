// Package goList 提供了对泛型切片的常见操作。
package goList

import "errors"

type slice[T any] struct {
	Slice  []T
	Length int
	zero   T
}

//type stack[T any] struct {
//	items []T
//	zero  T
//}

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

// 判断切片数组是否为空
func (s *slice[T]) isNullS() bool {
	if s.Length == 0 {
		return false
	}

	return true
}

// 检查索引是否处于合法范围内
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

// 将索引值转换为有效范围内的索引值
func (s *slice[T]) indexCon(index int) int {
	if index < 0 {
		index += s.Length
		if index < 0 {
			return 0
		}
	} else if index > s.Length {
		index = s.Length - 1
	}

	return index
}

//
//func stackNew[T any]() *stack[T] {
//	return &stack[T]{
//		items: make([]T, 0),
//	}
//}
//
//func (st *stack[T]) stackIsEmpty() bool {
//	return len(st.items) == 0
//}
//
//func (st *stack[T]) stackAdd(item T) {
//	st.items = append(st.items, item)
//}
//
//func (st *stack[T]) stackDel() (T, error) {
//	if st.stackIsEmpty() {
//		return st.zero, errors.New("stack is empty")
//	}
//
//	item := st.items[len(st.items)-1]
//	st.items = st.items[:len(st.items)-1]
//	return item, nil
//}
