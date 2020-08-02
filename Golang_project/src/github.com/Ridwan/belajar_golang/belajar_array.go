package main

import "fmt"

func main() {
	var a [7]int
	a[0] = 23
	a[1] = 70
	a[2] = 10
	a[3] = 45
	a[4] = 23
	a[5] = 25
	a[6] = 76
	fmt.Println(a)

	var buah [5]string
	buah[0] = "Semangka"
	buah[1] = "Apel"
	buah[2] = "Nanas"
	buah[3] = "Durian"
	buah[4] = "Pisang"
	fmt.Println(buah)
	fmt.Println(buah[3])

}
