package main

import "fmt"

func main() {
	for a := 0; a <= 5; a++ {
		// if a == 3 { berhenti di ke 3
		// 	break
		// }
		if a == 3 { //hanya memberhentika di no 3 saja lalu di lanjut
			continue
		}
		fmt.Println("Nilai ", a)
	}
}
