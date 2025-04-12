package tool

func RemoveDuplicates[T comparable](slice []T) []T {
	encountered := make(map[T]struct{})
	result := make([]T, 0, len(slice))
	for _, v := range slice {
		if _, ok := encountered[v]; !ok {
			encountered[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}
