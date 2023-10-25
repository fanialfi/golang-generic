package main

type UserMode[T int | float64] struct {
	Name   string
	Scores []T
}

func (m *UserMode[int]) SetScoresA(scores []int) {
	m.Scores = scores
}

func (m *UserMode[float64]) SetScoresB(scores []float64) {
	m.Scores = scores
}
