package main

// func Sum(number []int) int {
// 	var total int
// 	for _, elm := range number {
// 		total += elm
// 	}

// 	return total
// }

// generic
// penulisan notasi function pada generic
// func FuncName[dataType <ComparableType>](params)
func Sum[V int | float32 | float64](numbers []V) V {
	var total V
	for _, elm := range numbers {
		total += elm
	}

	return total
}
