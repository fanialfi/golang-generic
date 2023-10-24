package main

func SumNumberS1(m map[string]int64) int64 {
	var sum int64
	for _, value := range m {
		sum += value
	}
	return sum
}

func SumNumberS2[K comparable, V int64 | float64](m map[K]V) V {
	var sum V
	for _, value := range m {
		sum += value
	}

	return sum
}
