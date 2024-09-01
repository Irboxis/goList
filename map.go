package goList

import (
	"errors"
	"fmt"
)

// Map 对切片中的每个元素应用提供的回调函数，并返回一个新切片，包含所有满足回调函数条件的元素。
//
// 参数：
//   - callback: 接受三个参数的回调函数：
//     1. elem: 当前处理的切片元素。
//     2. index: 当前元素在切片中的索引。
//     3. slice: 当前正在处理的整个切片。
//     回调函数返回一个布尔值 `bool` 和一个值 `T`，如果布尔值为 `true`，则将值 `T` 添加到结果切片中；如果为 `false`，则返回错误并中止处理。
//
// 返回：
//   - 一个新切片，包含所有满足回调函数条件的元素。
//   - 如果切片为空，返回 `nil` 和错误。
func (s *slice[T]) Map(callback func(elem T, index int, slice []T) (bool, T)) ([]T, error) {
	if s.isNullS() {
		return nil, errors.New("currently an empty slice")
	}

	result := make([]T, 0)
	for index, val := range s.Slice {
		t, v := callback(val, index, s.Slice)
		if t {
			result = append(result, v)
		} else {
			return nil, errors.New(fmt.Sprintf("callback execution error, index is %d, value is %v", index, val))
		}
	}

	return result, nil
}

// Merge 接受一个或多个相同类型的数组或切片，并将它们与当前切片合并。
//
// 参数:
//   - list: 需要合并的一个或多个切片。
//
// 返回:
//   - 包含合并后结果的新切片。
func (s *slice[T]) Merge(list ...[]T) []T {
	totalLength := s.Length
	for _, l := range list {
		totalLength += len(l)
	}

	result := make([]T, 0, totalLength)
	result = append(result, s.Slice...)

	for _, l := range list {
		result = append(result, l...)
	}

	return result
}
