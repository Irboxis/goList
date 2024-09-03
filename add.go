package goList

// Add 从切片中的指定索引开始插入一个或多个元素。
//
// 索引可以是负值，在这种情况下，位置是相对于切片末尾计算的。如果索引超出范围（即小于 `-Length` 或大于 `Length`），则返回错误。
//
// 如果切片的容量不足以容纳新元素，则会自动调整大小。
//
// 参数：
//   - start：切片中将插入元素的位置，可以是负值。
//   - elements：从指定索引开始插入的元素的可变列表。
//
// 返回：
//   - 如果索引超出范围，则返回错误。否则，返回 nil。
func (s *list[T]) Add(start int, elements ...T) error {
	i, err := indexCheck(s, start)
	if err != nil {
		return err
	}

	if len(s.Slice)+len(elements) > cap(s.Slice) {
		newCap := len(s.Slice) + len(elements)
		newSlice := make([]T, len(s.Slice), newCap)
		copy(newSlice, s.Slice)
		s.Slice = newSlice
	}

	s.Slice = append(s.Slice[:i], append(elements, s.Slice[i:]...)...)
	s.Length += len(elements)

	return nil
}

// Push 在切片的末尾追加一个或多个元素。
//
// 该方法将指定的元素追加到切片的末尾。
// 如果追加的元素数量导致超过切片的当前容量，切片将自动扩容。
//
// 参数：
//   - elements：要追加到切片末尾的元素的可变列表。
func (s *list[T]) Push(elements ...T) {
	s.Slice = append(s.Slice, elements...)
}

// Shift 在切片的开头插入一个或多个元素。
//
// 该方法将指定的元素插入切片的起始位置，并将原有元素顺序后移。
// 如果插入的元素数量大于切片的当前容量，切片将自动扩容。
//
// 参数：
//   - elements：要插入到切片开头的元素的可变列表。
func (s *list[T]) Shift(elements ...T) {
	result := make([]T, len(elements), s.Length+len(elements))
	copy(result, elements)
	s.Slice = append(result, s.Slice...)
	s.Length += len(elements)
}
