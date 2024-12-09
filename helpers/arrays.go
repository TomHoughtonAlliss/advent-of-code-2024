package helpers

func GetPermutations[T comparable](in []T) [][]T {
	var result [][]T

	var permute func([]T, int)
	permute = func(a []T, k int) {
		if k == len(a) {
			perm := make([]T, len(a))
			copy(perm, a)
			result = append(result, perm)
		} else {
			for i := k; i < len(a); i++ {
				a[k], a[i] = a[i], a[k]
				permute(a, k+1)
				a[k], a[i] = a[i], a[k]
			}
		}
	}

	permute(in, 0)
	return result
}

func GetIntersection[T comparable](a []T, b []T) []T {
	set := make(map[T]any)
	for _, item := range a {
		set[item] = 0
	}

	var intersection []T
	for _, item := range b {
		if _, found := set[item]; found {
			intersection = append(intersection, item)
		}
	}

	return intersection
}

func Swap[T any](arr []T, i int, j int) []T {
	arr[i], arr[j] = arr[j], arr[i]
	return arr
}
