// ! Fungsi channel yaitu untuk menghubungkan go routine dengan go routine lainnya
// ! Tugas channel ini untuk mengirimkan dan menerima data
// ! channel ini juga bersifat ansycronous

package main

import "fmt"

func main() {
	// c := make(chan string)
	// go func(n string) {
	// 	c <- n
	// }("Ridwan")

	// var chanFunc = func(n string) {
	// 	c <- n
	// }

	// go chanFunc("AAAAA")

	// b := <-c
	// o := <-c
	// fmt.Println(b)
	// fmt.Println(o)

	// chanKedua := make(chan string)
	// go CetakNama(chanKedua, "Hahahhah")

	// var ll = <-chanKedua
	// fmt.Println(ll)

	// Buffered Channel
	pp := make(chan int, 2) //jadi ini untuk maksimal nya , di tambah 1
	go func() {
		for p := range pp {
			p = <-pp                               // si pp ini ngirim data ke si p
			fmt.Println("Terima Data ========", p) // jadi terima data nya maksimum tergantung pp (tidak akan melebihi)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("Kirim data")
		pp <- i
	}

}

func CetakNama(c chan string, v string) {
	c <- v
}
