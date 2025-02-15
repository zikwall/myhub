package x

func Map[T any, R any](collection []T, iteratee func(item T, index int) R) []R {
	result := make([]R, len(collection))
	for i, item := range collection {
		result[i] = iteratee(item, i)
	}
	return result
}

func RemoveElementByValue[T comparable](s []T, value T) []T {
	for i, v := range s {
		if v == value {
			// Если значение найдено, возвращаем новый срез без этого элемента.
			return append(s[:i], s[i+1:]...)
		}
	}
	// Если значение не найдено, возвращаем оригинальный срез.
	return s
}
