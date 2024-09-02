package goList

import (
	"errors"
	"fmt"
	"reflect"
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
	if !s.isNullS() {
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

// Every 方法遍历切片中的每个元素，并将其传递给指定的回调函数进行检查。
//
// 参数：
//   - callback：一个函数，该函数接受三个参数：
//     1. elem：当前正在处理的元素。
//     2. index：当前元素在切片中的索引。
//     3. slice：整个切片本身。
//
// 返回：
//   - bool：如果所有元素都满足回调函数的条件，则返回 true；否则返回 false。
func (s *slice[T]) Every(callback func(elem T, index int, slice []T) bool) bool {
	for i, v := range s.Slice {
		if !callback(v, i, s.Slice) {
			return false
		}
	}

	return true
}

// Fill 用于将切片中的指定范围填充为给定的元素。
//
// 参数:
//   - element: 要填充到切片中的元素。
//   - scope: 可选参数，指定填充的起始和结束索引范围；支持负索引模式。
//     1. 若为一个参数，则该参数为起始索引，结束索引默认为切片长度。
//     2. 若为两个参数，则第一个参数为起始索引，第二个参数为结束索引。
func (s *slice[T]) Fill(element T, scope ...int) error {
	if !s.isNullS() {
		return errors.New("currently an empty slice")
	}

	start, end := 0, s.Length

	if len(scope) == 1 {
		start = s.indexCon(scope[0])
	} else if len(scope) == 2 {
		start = s.indexCon(scope[0])
		end = s.indexCon(scope[1])
	} else if len(scope) > 2 {
		return errors.New("unexpected number of parameters, scope can only accept two parameters as an index range")
	}

	if start >= end {
		return errors.New("unexpected scope parameter, start index should be less than end index")
	}

	for i := start; i < end; i++ {
		s.Slice[i] = element
	}

	return nil
}

// Filter 遍历切片中的每个元素，并将满足条件的元素返回为一个新的切片。
//
// 参数:
//   - callback: 一个回调函数，接受切片中的每个元素、其索引以及整个切片作为参数。
//
// 返回:
//   - 包含所有满足条件元素的切片。
//   - 如果切片为空，则返回 nil 和一个错误。
func (s *slice[T]) Filter(callback func(elem T, index int, slice []T) bool) ([]T, error) {
	if !s.isNullS() {
		return nil, errors.New("currently an empty slice")
	}

	result := make([]T, 0, s.Length)
	for index, val := range s.Slice {
		if callback(val, index, s.Slice) {
			result = append(result, val)
		}
	}

	return result, nil
}

func (s *slice[T]) Flat(deep ...int) ([]T, error) {
	if !s.isNullS() {
		return nil, errors.New("currently an empty slice")
	}

	deepInit := 1
	if len(deep) == 1 {
		deepInit = deep[0]
	} else if len(deep) > 1 {
		return nil, errors.New("unexpected number of parameters, deep only accepts one parameter as the flattening depth")
	}

	result := make([]T, 0)
	stack := []any{s.Slice}
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		val := reflect.ValueOf(current)
		if val.Kind() == reflect.Slice || val.Kind() == reflect.Array {
			if deepInit > 0 {
				for i := val.Len() - 1; i >= 0; i-- {
					stack = append(stack, val.Index(i).Interface())
				}
				deepInit--
			} else if deepInit == -1 {
				for i := val.Len() - 1; i >= 0; i-- {
					stack = append(stack, val.Index(i).Interface())
				}
			} else {
				// 深度为 0 时，将当前层次的元素添加到结果中
				for i := 0; i < val.Len(); i++ {
					result = append(result, val.Index(i).Interface().(T))
				}
			}
		} else {
			result = append(result, current.(T))
		}
	}

	return result, nil
}
