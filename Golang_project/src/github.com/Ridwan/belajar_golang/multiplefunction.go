package main

import "fmt"

func main() {
	a := CetakNama("Muhamad", "Ridwan")
	fmt.Println(a)

	b, c, d, e := Calculator(2, 8, 10)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}

func CetakNama(namaDepan, namaBelakang string) string {
	fullName := namaDepan + " " + namaBelakang
	return fullName
}

func Calculator(nilai1, nilai2, nilai3 int) (int, int, int, float64) { //jadi ketika kita me return lebih dari 2 maka tipe data nya juga harus 2 dan memakai kurung jika lebih dari  1 tipe datanya
	pertambahan := nilai1 + nilai2 + nilai3
	pengurangan := nilai1 - nilai2 - nilai3
	perkalian := nilai1 * nilai2 * nilai3
	pembagian := float64(nilai1) / float64(nilai2) / float64(nilai3)

	return pertambahan, pengurangan, perkalian, pembagian
}
