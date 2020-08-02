package main

import (
	"basic_golang/calculator"
	"basic_golang/persegi"
	"fmt"

)

func main() {
	fmt.Println("Test")

	var c calculator.HitungCalculator
	c = calculator.Nilai{4, 7, 10} //sesuaikan dengan nama packagenya
	fmt.Println(c.Perkalian())

	var p persegi.HitungPersegi
	p = persegi.Persegi{90}
	fmt.Println(p.Luas())
}
