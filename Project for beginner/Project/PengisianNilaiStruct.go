/* File: PengisianNilaiStruct.go*/
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Pengisian sebuah nilai struct dapat dilakukan dengan cara mengakses properti struct
terlebih dahulu. Pengaksesan properti pada struct menggunakan . (dot) yang diikuti oleh nama
instance dan nama properti dari struct tersebut, contoh:
*/
/*---------------*/
package main

import "fmt"

type Student struct {
	Matric_number string
	Name          string
	Age           uint
}

func main() {
	var arsyad Student //membuat instance baru
	arsyad.Matric_number = "12345"
	arsyad.Name = "gusti arsyad"
	arsyad.Age = 25
	fmt.Printf("Student: %+v\n", arsyad)
}
