// ! adalah sebuah function yang bisa disimpan di dalam variable
// ! closure function juga bisa di implementasikan function didalam function
//  ! bisa juga disebun anynomous function atau function tanpa nama

package main

import "fmt"

func main() {
	nama := func(n string) string {
		return n
	}

	a := nama("Muhamad Ridwan")

	fmt.Println(a)
	fmt.Println(nama("Muhamad Ridwan"))

	func(n string) {
		fmt.Println(n)
	}("Backend Programmer")

	func(a, b int) {
		fmt.Println("Hasil Pertambahan :", a+b)
	}(10, 5)

	pert := func(a, b int) int {
		return a + b
	}

	fmt.Println(pert(9, 8))

	b := FullName()
	fmt.Println(b("wkwkwkwk"))

	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	result := GanjilGenap() //result ini akan menjadi function
	gan, gen := result(num...)
	fmt.Println("Hasil Ganjil : ", gan)
	fmt.Println("Hasil Genap : ", gen)

}

func FullName() func(string) string {
	return func(n string) string {
		return n
	}
}

func GanjilGenap() func(...int) ([]int, []int) { //functon closure di mix dengan variadic function dimana di mix lagi dengan multiple return
	return func(n ...int) ([]int, []int) {
		genap, ganjil := []int{}, []int{}

		for _, val := range n {
			if val%2 == 0 {
				genap = append(genap, val)
			} else {
				ganjil = append(ganjil, val)
			}
		}

		return ganjil, genap

	}
}
