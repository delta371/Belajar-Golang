// ! adalah sebuah method atau function yang dapat di eksekusi secara independen dengan Gorountine lainnya
// ! untuk membuat Gorountine ini di tandai dengan keyword Go
// ! jadi Gorountine ini bersifat asyncronous

// ! Fungsi defer yaitu funngsi untuk mengakhiri eksekusi sebuah statement sebelum si code itu selesai
// ! jadi jika memakai fungsi defer maka ia akan di eksekusi di akhir (dihitung mundur)

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	// defer fmt.Println("Hahahaha")
	// defer fmt.Println("Hoooofofo")
	// fmt.Println("Hihihih")

	// for i := 0; i < 5; i++ {
	// 	defer fmt.Println(i)
	// }

	wg.Add(3) //untuk memberi tahu jumlah go rountine itu sendiri
	go CetakNama("Hi")
	go CetakNama("Hello")
	go CetakNama("What's Up Bro")
	wg.Wait()
}

func CetakNama(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100) //bebas di pakai atau tidak
	}

	defer wg.Done() // memberi tahu syncronize selesai
}
