package main

import "fmt"

func main() {
	var buah map[string]int //deklarasi
	buah = map[string]int{} //inisialisasi

	buah["semangka"] = 20000
	buah["apel"] = 5000

	fmt.Println(buah)
	fmt.Println("Harga buah apel : ", buah["apel"])

	delete(buah, "apel") //fungsi dari golang untuk mendownload item, diikuti dengan nama variable dan key nya
	fmt.Println("Final Result : ", buah)

	a := map[int]int{}
	a[0] = 12
	a[1] = 13
	fmt.Println(a)

	b := make(map[int]int)
	b[10] = 1000
	b[11] = 2000
	b[12] = 3000
	b[13] = 4000

	delete(b, 12)

	fmt.Println(b[12])

	value, isB := b[12] //isB adalah untuk mereturn valuie boolean
	fmt.Println("Value : ", value)
	fmt.Println("Cek b 12", isB)

}
