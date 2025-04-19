package tool

// SliceCopy 从切片中安全地复制元素
// 参数:
//   - slice: 源切片
//   - start: 开始位置(包含)
//   - count: 要取的元素数量
//
// 返回:
//   - 新切片，包含从start开始的count个元素
//   - 如果start超出范围，返回空切片
//   - 如果count超出剩余元素数量，返回剩余所有元素
func SliceCopy[T any](slice []T, start, count int) []T {
	// 处理start超出范围的情况
	if start < 0 || start >= len(slice) {
		return []T{}
	}

	// 计算实际可以复制的数量
	available := len(slice) - start
	copyCount := min(count, available)

	// 创建新切片并复制
	result := make([]T, copyCount)
	copy(result, slice[start:start+copyCount])

	return result
}

// FirstN 获取前N个元素的副本
func FirstN[T any](slice []T, n int) []T {
	return SliceCopy(slice, 0, n)
}

// LastN 获取最后N个元素的副本
func LastN[T any](slice []T, n int) []T {
	start := len(slice) - n
	if start < 0 {
		start = 0
	}
	return SliceCopy(slice, start, n)
}

// FromIndex 从指定索引开始到结尾的所有元素
func FromIndex[T any](slice []T, start int) []T {
	return SliceCopy(slice, start, len(slice)-start)
}

// Between 获取从start到end(不包含)之间的元素
func Between[T any](slice []T, start, end int) []T {
	if start < 0 || end <= start || start >= len(slice) {
		return nil
	}
	return SliceCopy(slice, start, end-start)
}
