package goList

import (
	"errors"
	"reflect"
)

// IndexOf 在切片中查找指定元素第一次出现时的索引。支持结构体、函数等复杂类型的查找。
//
// 参数：
//   - element：要查找的元素。
//
// 返回：
//   - 如果找到匹配的元素，返回其在切片中的索引。
//   - 如果未找到匹配的元素，返回 -1 和错误，指示元素未找到。
func (s *list[T]) IndexOf(element T) (int, error) {
	if !s.isNullS() {
		return -1, errors.New("currently an empty slice")
	}

	for index, val := range s.Slice {
		if reflect.DeepEqual(val, element) {
			return index, nil
		}
	}

	return -1, errors.New("index not found")
}

// ValueOf 返回切片中指定索引处的元素
//
// 如果 index 超出切片范围,小于 0 的情况下将重置 index 为 0 , 大于切片最大长度的情况下将重置为最后一位元素的索引
//
// 参数：
//   - index：要获取元素的位置，可以是负值，表示从切片末尾开始计算
//
// 返回：
//   - 在指定索引处的元素。
//   - 如果切片为空或元素未找到，则返回切片类型的零值和相应的错误
func (s *list[T]) ValueOf(index int) (T, error) {
	if !s.isNullS() {
		return s.zero, errors.New("currently an empty slice")
	}
	index = s.indexCon(index)

	for i, val := range s.Slice {
		if i == index {
			return val, nil
		}
	}

	return s.zero, errors.New("element not found")
}

// Find 遍历切片的元素并将提供的回调函数应用于每个元素。
//
// 参数：
//   - 回调：以切片的元素、其索引和整个切片作为参数的函数。
//
// 返回：
//   - 满足回调条件的第一个元素，或者如果切片为空或没有元素符合条件，则返回元素类型的零值。
//   - 错误，指示切片是否为空或未找到匹配元素。
func (s *list[T]) Find(callback func(elem T, index int, slice []T) bool) (T, error) {
	if !s.isNullS() {
		return s.zero, errors.New("currently an empty slice")
	}

	for index, val := range s.Slice {
		if callback(val, index, s.Slice) {
			return val, nil
		}
	}

	return s.zero, errors.New("element not found")
}

// Includes 函数用于检查切片中是否存在指定的元素，搜索从给定的索引开始（如果未提供起始索引，默认从第一个元素开始）。
//
// 参数:
//   - element: 要在切片中搜索的元素。
//   - start: （可选）开始搜索的索引。
//
// 返回:
//   - bool: 如果找到了该元素，则返回 true，否则返回 false。
//   - error: 如果起始索引超出范围或提供了多个起始值，则返回错误信息。
func (s *list[T]) Includes(element T, start ...int) (bool, error) {
	if !s.isNullS() {
		return false, errors.New("currently an empty slice")
	}
	index := 0
	if len(start) == 1 {
		index = s.indexCon(start[0])
	} else if len(start) > 1 {
		return false, errors.New("unexpected argument, start only accepts one value as the starting index")
	}

	for i := index; i < s.Length; i++ {
		if reflect.DeepEqual(element, s.Slice[i]) {
			return true, nil
		}
	}

	return false, nil
}
