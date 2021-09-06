/* File: ValueStruct.go*/
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Jika urutan properti sudah diketahui, nama properti dapat dihilangkan pada saat
pendeklarasian struct, contoh:
"Note: Urutan properti perlu diperhatikan pada saat menggunakan cara di atas, dikarenakan
Golang akan memasukkan nilai ke dalam properti sesuai dengan urutan, "
*/
/*---------------*/
package main

import "fmt"

type Student struct {
	Matric_number string
	Name          string
	age           uint
}

func main() {
	arsyad := Student{
		"12345",
		"gusti arsyad",
		25,
	}
	fmt.Printf("Student: %+v\n", arsyad)

}
