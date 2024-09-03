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
func (s *list[T]) Update(callback func(elem T, index int, slice *[]T) bool) error {
	for index, val := range s.Slice {
		if !callback(val, index, &s.Slice) {
			return errors.New(fmt.Sprintf("unexpected error: index is %d, elem is %v", index, val))
		}
	}

	return nil
}
