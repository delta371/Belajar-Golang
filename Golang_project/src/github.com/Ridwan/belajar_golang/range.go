// ! untuk mengulang sebuah kode
// ! biasa nya ini di pakai untuk data yang bertipe array,slice dan map
// ! perbedaan range dengan forloop yaitu range menggunakan balikan 2 data (index dan element)

package main

import "fmt"

func main() {
	a := []string{"a", "b", "c", "d", "e"}
	fmt.Println(a)

	fmt.Println("==========================================")

	for i := 0; i < len(a); i++ {
		fmt.Println(i)
	}

	fmt.Println("==========================================")

	for _, value := range a { // arti (_) dia tetap mengembalikkan 2 data tapi di data yang pertama (_) saya tidak gunakan atau abaikan
		// fmt.Println("index menggunakan range: ", index)
		fmt.Println("index menggunakan value: ", value)
	}

	fmt.Println("==========================================")

	// OR

	for index := range a {
		fmt.Println("index menggunakan range2 : ", a[index])
	}

	fmt.Println("==========================================")

	b := map[string]int{"mobil": 1000, "motor": 2000, "hanphone": 3000}
	fmt.Println(b)

	for i, v := range b {
		fmt.Println("index: ", i)
		fmt.Println("value: ", v)
	}

	fmt.Println("==========================================")

	c := map[int]int{0: 1000, 1: 2000, 2: 3000}
	for v := range c {
		fmt.Println("value: ", v)
	}

}
