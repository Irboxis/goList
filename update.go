package goList

import (
	"errors"
	"fmt"
)

// Update 根据回调函数更新当前切片中的每个元素。
//
// 对切片中的每个元素调用回调函数，并允许修改切片中的元素。如果回调函数返回 false，更新将停止，并返回一个错误。
//
// 参数：
//   - callback：一个回调函数，该函数接受当前元素、索引和切片指针，返回一个布尔值表示是否继续更新。
//
// 返回：
//   - 如果回调函数返回 false，返回一个错误信息。否则，返回 nil。
func (s *slice[T]) Update(callback func(elem T, index int, slice *[]T) bool) error {
	for index, val := range s.Slice {
		if !callback(val, index, &s.Slice) {
			return errors.New(fmt.Sprintf("unexpected error: index is %d, elem is %v", index, val))
		}
	}

	return nil
}

// ToUpdate 根据回调函数更新并返回一个新的切片。
//
// 对切片中的每个元素调用回调函数，并在返回的新切片中应用更新。如果回调函数返回 false，更新将停止，并返回一个错误。
//
// 参数：
//   - callback：一个回调函数，该函数接受当前元素、索引和切片，返回一个布尔值表示是否继续更新。
//
// 返回：
//   - 一个更新后的新切片。
//   - 如果回调函数返回 false，返回一个错误信息。否则，返回更新后的切片和 nil。
func (s *slice[T]) ToUpdate(callback func(elem T, index int, slice []T) bool) ([]T, error) {
	result := make([]T, 0, s.Length)
	copy(result, s.Slice)

	for index, val := range result {
		if !callback(val, index, result) {
			return nil, errors.New(fmt.Sprintf("unexpected error: index is %d, elem is %v", index, val))
		}
	}

	return result, nil
}
