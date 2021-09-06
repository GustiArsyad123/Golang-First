/* File: varTemStatementIf-Else.go*/
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Pengertian dari variabel temporary adalah variabel yang hanya bisa digunakan pada
blok seleksi kondisi dimana ia ditempatkan saja. Penggunaan variabel ini membawa beberapa
manfaat, antara lain:
	1. Scope atau cakupan variabel jelas, hanya bisa digunakan pada blok seleksi kondisi
	itu saja.
	2. Kode menjadi lebih rapih.
	3. Ketika nilai variabel tersebut didapat dari sebuah komputasi, perhitungan tidak
	perlu dilakukan di dalam blok masing-masing kondisi.
	Bentuk umumnya sama seperti statement if-else biasa, yang berbeda adalah penggunaan dari
	nilai variabelnya saja
*/
/*---------------*/
package main

import "fmt"

func main() {
	var value = 9045.0
	if percent := value / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect\n", percent, "%")
	} else if percent >= 80 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}
}
