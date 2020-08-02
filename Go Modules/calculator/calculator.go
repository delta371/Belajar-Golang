package calculator

type HitungCalculator interface {
	Perkalian() int
}

type Nilai struct {
	A, B, C int
}

func (n Nilai) Perkalian() int {
	return n.A * n.B * n.C
}
