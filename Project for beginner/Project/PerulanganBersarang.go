/* File: PerulanganBersarang.go */
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Cara pengaplikasian perulangan bersarang pada pemrograman Go kurang lebih sama
dengan perulangan bersarang pada bahasa pemrograman lainnya, yaitu dengan menulikan blok
statement perulangan didalam perulangan.
bentuk umum:
For inisialisasi;kondisi;iterasi{
	For inisialisasi;kondisi;iterasi
		{ statement }
	statement
}
*/
/*---------------*/
package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		for j := i; j < 10; j++ {
			fmt.Print(j, " ")
		}
	}
	fmt.Println()
}
