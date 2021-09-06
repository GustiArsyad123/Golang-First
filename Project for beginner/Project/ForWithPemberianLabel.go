/* File: ForWithPermberianLabel.go */
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Pada salah satu kasus perulangan bersarang, break dan continue akan berlaku pada blok
perulangan dimana break dan continue jika digunakan. Ada cara agar kedua kata kunci ini bisa
tertuju pada perulangan terluar atau perulangan tertentu, yaitu dengan memanfaatkan teknik
pemberian label.
Bentuk umum :
nama_label:
For inisialisasi;kondisi;iterasi{
	For inisialisasi;kondisi;iterasi{
		kondisi {
		break nama_label
		}
	statement
	}
}
*/
/*---------------*/
package main

import "fmt"

func main() {
LoopOuter:
	for i := 0; i < 10; i++ {
		for j := i; j < 10; j++ {
			for i == 2 {
				break LoopOuter
			}
			fmt.Print("Matrix [", i, "][", j, "]", "\n")
		}
	}
}
