// ! adalah sebuah function yang memiliki reciver argument
// ! reciver ini bisa itu bis sebuah type struct atau none type struct

package main

import "fmt"

type Nilai struct {
	A, B int
}

type Angka struct {
	A, B int
}

func main() {
	result := Nilai{2, 4}
	fmt.Println(result.Pertambahan())
	// result.Pertambahan1()
	fmt.Println(result.Perkalian())

	result.UbahNilaiA(15)
	fmt.Println(result.A)

	fmt.Println("Hasil pertambahan nilai a yang baru : ", result.Pertambahan())

	l := Angka{22, 33}
	l.Pertambahan()
}

func (n Nilai) Pertambahan() int { //Method
	return n.A + n.B
}

func (n Nilai) Perkalian() int { //Method
	return n.A * n.B
}

// func (n Nilai) Pertambahan1() {
// 	fmt.Println("Hasil method bukan return value :", n.A+n.B)
// }
func (n Angka) Pertambahan() { //Method
	fmt.Println("Hasil method bukan return value :", n.A+n.B)
}

func (n *Nilai) UbahNilaiA(a int) {
	n.A = a
}
