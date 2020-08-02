//  ! variadicfunction adalah sebuah function yang memiliki unlimited parameter/memiliki parameter yang tak terhingga

package main

import "fmt"

func main() {
	TesVariadic(1)
	TesFunction(1)
	p := Pertambahan(2, 3, 4, 5)
	fmt.Println("Hasil Pertambahan dengan sebelumnya : ", p)

	s := []int{2, 4, 5, 8, 10}
	p2 := Pertambahan2(s...)
	fmt.Println("Hasil Pertambahan dengan sebelumnya Versi slice : ", p2)
}

func TesVariadic(number ...int) { // Hasilnya akan berbentuk array , dan bisa di tambahkan banyak parameter
	fmt.Println(number)
}

func TesFunction(number int) {
	fmt.Println(number)
}

func Pertambahan(n ...int) int {
	result := 0

	for i := 0; i < len(n); i++ {
		result += n[i]
	}
	return result
}

func Pertambahan2(n ...int) int {
	result := 0

	for i := 0; i < len(n); i++ {
		result += n[i]
	}
	return result
}
