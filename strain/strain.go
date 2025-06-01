package strain

// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1

func Keep[T any](list []T, fn func(T) bool) []T {
	var result []T
	for _, i := range list {
		if fn(i) {
			result = append(result, i)
		}
	}
	return result
}

func Discard[T any](list []T, fn func(T) bool) []T {
	var result []T
	for _, i := range list {
		if !fn(i) {
			result = append(result, i)
		}
	}
	return result
}
