package main

import "fmt"

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

func main() {
	// cara pertama pemanggilan function generic
	total1 := Sum([]int{1, 2, 3, 4, 5})
	total2 := Sum([]float32{2.5, 7.2})
	total3 := Sum([]float64{1.23, 6.33, 12.6})

	fmt.Printf("total : %d\n", total1)
	fmt.Printf("total : %.3f\n", total2)
	fmt.Printf("total : %.3f\n", total3)

	fmt.Println()
	// cara kedua pemanggilan function generic
	total4 := Sum[int]([]int{1, 2, 3, 4, 5})
	total5 := Sum[float32]([]float32{2.5, 7.2})
	total6 := Sum[float64]([]float64{1.23, 6.33, 12.6})

	fmt.Printf("total : %d\n", total4)
	fmt.Printf("total : %.3f\n", total5)
	fmt.Printf("total : %.3f\n", total6)
}
