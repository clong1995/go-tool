package tool

func Contains[T comparable](slice []T, item T) (exists bool) {
	for _, v := range slice {
		if v == item {
			exists = true
			return
		}
	}
	return
}
