package util

func ContainsAtIndex[V int | rune | string](slice []V, value V) int {
	for idx, element := range slice {
		if element == value {
			return idx
		}
	}
	return -1
}

func Contains[V int | rune | string](slice []V, value V) bool {
	return ContainsAtIndex[V](slice, value) != -1
}
