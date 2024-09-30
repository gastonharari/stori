package structs

func Map[T1 any, T2 any](input []T1, fn func(T1) T2) []T2 {
	output := make([]T2, 0, len(input))
	for _, item := range input {
		output = append(output, fn(item))
	}
	return output
}
