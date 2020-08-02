// ! adalah sebuah kumpulan dari definisi method yang tidak memiliki isi dan di bungkus dalam nama tertentu
// ! bisa juga di pakai sebagai datatype atau tipe data
// ! Kumpulan dari method"

package main

import (
	"fmt"
	"math"
)

type Calculator interface {
	Pertambahan() int
	Perkalian() int
}

type HitungPersegi interface {
	Keliling() int
	Luas() int
}

type Nilai struct {
	A, B int
}

type Persegi struct {
	Sisi int
}

func main() {
	var c Calculator
	c = Nilai{10, 20}
	fmt.Println("Hasil nilai Pertambahan : ", c.Pertambahan())
	fmt.Println("Hasil nilai Perkalian : ", c.Perkalian())

	var hp HitungPersegi
	hp = Persegi{9}
	fmt.Println("Hasil Keliling Persegi : ", hp.Keliling())
	fmt.Println("Hasil Luas Persegi : ", hp.Luas())
}

func (n Nilai) Pertambahan() int {
	return n.A + n.B
}

func (n Nilai) Perkalian() int {
	return n.A * n.B
}

func (s Persegi) Keliling() int {
	return 4 * s.Sisi
}

func (s Persegi) Luas() int {
	// return s.Sisi * s.Sisi
	return int(math.Pow(float64(s.Sisi), 2)) //meng convert ke int lagi
}
