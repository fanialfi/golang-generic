package main

import "fmt"

// struct generic
func main() {
	var m1 UserMode[int]
	m1.Name = "fani"
	m1.Scores = []int{1, 2, 3}
	fmt.Printf("score %#v\n", m1)

	var m2 UserMode[float64]
	m2.Name = "fani"
	m2.SetScoresB([]float64{10, 11})
	fmt.Printf("score %#v\n", m2)
}

// func main() {
// 	ints := map[string]int64{"first": 34, "second": 12}
// 	floats := map[string]float64{"first": 35.98, "second": 26.99}
// 	fmt.Printf("Generic Sum With Constrain : %v and %v\n",
// 		SumNumberS2[string](ints),
// 		SumNumberS2[string](floats),
// 	)
// }

// func main() {
// 	// cara pertama pemanggilan function generic
// 	total1 := Sum([]int{1, 2, 3, 4, 5})
// 	total2 := Sum([]float32{2.5, 7.2})
// 	total3 := Sum([]float64{1.23, 6.33, 12.6})

// 	fmt.Printf("total : %d\n", total1)
// 	fmt.Printf("total : %.3f\n", total2)
// 	fmt.Printf("total : %.3f\n", total3)

// 	fmt.Println()
// 	// cara kedua pemanggilan function generic
// 	total4 := Sum[int]([]int{1, 2, 3, 4, 5})
// 	total5 := Sum[float32]([]float32{2.5, 7.2})
// 	total6 := Sum[float64]([]float64{1.23, 6.33, 12.6})

// 	fmt.Printf("total : %d\n", total4)
// 	fmt.Printf("total : %.3f\n", total5)
// 	fmt.Printf("total : %.3f\n", total6)
// }
