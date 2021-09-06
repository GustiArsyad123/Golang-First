/* File: MultiVariable.go */
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Golang dalam mendeklarasikan variabel secara
bersamaan. Deklarasi ini dapat dituliskan dalam satu baris perintah deklarasi variabel dan menggunakan
tanda koma ( , ) sebagai pembatas.
Bentuk umum 1:
var nama_var1, nama_var2, nama_varN tipe_data
nama_var1, nama_var2, nama_varN = nilai1, nilai2, nilaiN
Bentuk umum 2:
var nama_var1, nama_var2, nama_varN tipe_data = nilai1, nilai2, nilaiN
Bentuk umum 3:
nama_var1, nama_var2, nama_varN := nilai1, nilai2, nilaiN
*/
/*---------------*/
package main

import "fmt"

func main() {
	var a1, a2, a3 string
	a1, a2, a3 = "I", "Love", "Indonesia"

	var b1, b2, b3 float64 = 1, 20.7, 10

	c1, c2, c3 := "Hello, I'm Arsyad", false, 0

	/*Printah untuk menampilkan nilai dari sebuah variable*/
	fmt.Println(a1, a2, a3)
	fmt.Println(b1, b2, b3)
	fmt.Println(c1, c2, c3)

}
