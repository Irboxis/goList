package goList

import "errors"

// Delete 从切片中指定的索引处开始删除一个或多个元素。
//
// 参数：
//   - start：从哪个索引开始删除元素，可以是负值。
//   - count：要删除的元素数量，默认为1。
//
// 返回：
//   - 一个包含已删除元素的切片，以及可能的错误。
func (s *slice[T]) Delete(start int, count ...int) ([]T, error) {
	// delete an element by default
	numToDelete := 1
	if len(count) > 1 {
		return nil, errors.New("unexpected parameter, count can only accept one value as the number of elements to be deleted")
	} else if len(count) == 1 {
		numToDelete = count[0]
	}
	if numToDelete < 0 {
		return nil, errors.New("unexpected parameter, count cannot be negative")
	}

	i, err := indexCheck(s, start)
	if err != nil {
		return nil, err
	}

	if i+numToDelete > s.Length {
		numToDelete = s.Length - i
	}

	result := make([]T, numToDelete)
	copy(result, s.Slice[i:i+numToDelete])

	s.Slice = append(s.Slice[:i], s.Slice[i+numToDelete:]...)
	s.Length -= numToDelete

	return result, nil
}

// Pop 从切片末尾删除并返回最后一个元素。
//
// 返回：
//   - 被删除的元素。
func (s *slice[T]) Pop() (T, error) {
	if !s.isNullS() {
		return s.zero, errors.New("there are no elements in the slice")
	}

	result := s.Slice[s.Length-1]
	s.Slice = s.Slice[:s.Length-1]
	s.Length -= 1

	return result, nil
}

// UnShift 从切片开头删除并返回第一个元素。
//
// 返回：
//   - 被删除的元素。
func (s *slice[T]) UnShift() (T, error) {
	if !s.isNullS() {
		return s.zero, errors.New("there are no elements in the slice")
	}
	result := s.Slice[0]
	s.Slice = s.Slice[1:]
	s.Length -= 1

	return result, nil
}
