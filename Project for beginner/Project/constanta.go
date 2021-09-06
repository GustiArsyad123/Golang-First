/* File: constanta.go */
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Konstanta adalah variabel yang memiliki nilai tetap atau tidak bisa diubah. Inisialisasi nilai
hanya dilakukan sekali pada awal program, setelah itu data tidak dapat diubah nilainya.

Note:
Penggunaan Fungsi fmt.Print()
Fungsi ini memiliki peran yang sama seperti fungsi fmt.Println(), pembedanya fungsi
fmt.Print() tidak menghasilkan baris baru di akhir hasil keluarannya.
*/
/*---------------*/
package main

import "fmt"

func main() {
	const phi float64 = 3.14
	var jari, luas float64

	fmt.Print("inputkan nilai jari-jari = ")

	fmt.Scan(&jari)

	luas = phi * (jari * jari)
	fmt.Println("hasil luas lingkaran = ", luas)
}
