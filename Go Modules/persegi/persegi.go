package persegi

type HitungPersegi interface {
	Luas() int
}

type Persegi struct {
	Sisi int
}

func (p Persegi) Luas() int {
	return p.Sisi * p.Sisi
}
