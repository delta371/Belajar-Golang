// ! adalah referense atau alamat memori
// ! setiap variable nya adalah memori location
//  ! dan setiap memory location itu dapat di akses dengan menggunakan ampersan operator

package main

import "fmt"

func main() {
	var a int = 10
	fmt.Println(a)
	fmt.Println(&a) //yang disebut ampersan value

	// untuk mendeklarasikan pointer
	var b *int
	b = &a

	fmt.Println("Hasil alamat : ", b)
	fmt.Println("Hasil Value : ", *b) //Cara mengetahui value si b

	// untuk menggunakan pointer ini ga semua di variable aja tapi bisa juga di gunakan di function di parameternya

	aa := 5
	bb := 4
	Calc(&aa, &bb)

	fmt.Println("=====================")
	c := &aa
	d := &c
	e := &d

	fmt.Println("Value dari alamat c : ", d)
	fmt.Println("Value dari alamat si aa : ", c)
	fmt.Println("Value aa : ", aa)
	fmt.Println("Value dari pointer c si aa : ", *c)
	fmt.Println("Value dari pointer d : ", *d)
	fmt.Println("Value dari pointer ponter d : ", **d)
	fmt.Println("Value dari pointer ponter pointer e : ", ***e)

}

func Calc(a, b *int) {
	fmt.Println(*a + *b)
}
