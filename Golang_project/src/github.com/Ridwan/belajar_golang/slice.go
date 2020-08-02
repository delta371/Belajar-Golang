package main

import "fmt"

func main() {
	var array [3]int //array
	array[0] = 24

	array1 := [3]int{25} //array

	slice := []int{99, 77, 22, 33}
	slice2 := make([]int, 3)

	fmt.Println(array)
	fmt.Println(array1)
	fmt.Println(slice)
	fmt.Println(slice2)

	fmt.Println("Panjang array", len(array))
	fmt.Println("Panjang array1", len(array1))
	fmt.Println("Panjang slice", len(slice))

	slice = append(slice, 1000)
	slice = append(slice, 37842)
	slice = append(slice, 12110)
	slice = append(slice, 1312)
	fmt.Println(slice)

	slice2[0] = 20
	slice2[1] = 21
	slice2 = append(slice2, 25) // jadi bisa di tambah walaupu di awalnya di declarasikan
	slice2 = append(slice2, 26)
	fmt.Println(slice2)

	fmt.Println(slice2[3])

	for i := 0; i < len(slice2); i++ { //utuk mempermudah kita dalam pemanggilan array
		if slice2[i]%2 == 0 {
			fmt.Println("Hasil Perulangan Genap", slice2[i])
		} else {
			fmt.Println("Hasil Perulangan Ganjil", slice2[i])
		}
	}

}
