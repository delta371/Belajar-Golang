//  ! adalah sebuah kumpulan dari variable atau properti atau function yang dibungkus sebagai tipe data baru

package main

import "fmt"

type User struct {
	FirstName string
	LastName  string
	Addres    string
	Age       int
}

type Customer struct {
	JenisKelamin string
	User         User //diambil dari tipe data yang di atas
}

func main() {
	// u := User{"Muhamad", "Ridwan", "Dampit", 22}
	// u := User{
	// 	FirstName: "Muhamad",
	// 	LastName:  "Ridwan",
	// 	Addres:    "Dampit",
	// 	Age:       22,
	// }

	u := User{}

	u.FirstName = "Muhamad"
	u.LastName = "Ridwan"
	u.Addres = "Dampit"
	u.Age = 22
	fmt.Println(u)

	fmt.Println("Nama belakang : ", u.LastName)
	fmt.Println("Umur : ", u.Age)

	fmt.Println("===========================")

	// u di tampung didalam sebuah variable berbentuk pointer
	u1 := &u
	fmt.Println(u1)
	fmt.Println(*u1)
	fmt.Println(u1.FirstName)
	fmt.Println(u1.LastName)

	c := Customer{
		JenisKelamin: "Laki-Laki",
		User: User{
			FirstName: "Muhamad",
			LastName:  "Ridwan",
			Addres:    "Bogor",
			Age:       22,
		},
	}
	fmt.Println(c)
}
