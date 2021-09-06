/* File: EmbeddedStruct.go*/
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Embedded struct adalah sebuah struct yang diletakkan di dalam sebuah struct lain.
Struct jenis ini merepresentasikan sebuah hubungan pada struct, contoh:

Pada contoh code dibawah ini, struct Mahasiswa memiliki relasi one-to-one terhadap struct
StudentSite. Karena satu mahasiswa memiliki satu akun studentsite. Untuk mendefinisikan
bahwa kedua struct tersebut saling berhubungan diperlukan sebuah field yang menjembatani
kedua struct tersebut, dalam hal ini field AccountStudentSite pada struct Mahasiswa.
*/
/*---------------*/
package main

import "fmt"

type Student struct {
	Matric_number  string
	Name           string
	Age            uint
	AccountStudent StudentSite
}

type StudentSite struct {
	username string
	password string
}

func main() {
	var arsyad Student
	fmt.Printf("Student: %+v\n", arsyad)
}
