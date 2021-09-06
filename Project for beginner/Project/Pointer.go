/* File: Pointer.go*/
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Pointer adalah sebuah variabel yang nilainya berisi alamat memori dari variabel lain.
Dalam bahasa pemrograman Golang pendeklarasian sebuah pointer ditandai dengan * (asterik)
yang diikuti oleh tipe data dari variabel tersebut,

Pada Golang untuk mengambil alamat memori dari sebuah variabel dapat
menggunakan tanda & (ampersand) yang diletakkan sebelum nama dari variabel (dalam
notasi heksadesimal)

Sedangkan untuk mengambil atau mengubah nilai dari pointer dapat menggunakan
tanda * (asterik) yang diletakkan sebelum nama dari variabel,
*/
/*---------------*/
package main

import "fmt"

func main() {
	var x int
	var y *int //pointer declaration

	x = 10
	y = &x //y adalah alamat memori dari x.

	*y = 5

	fmt.Println("value x =", x)
	fmt.Println("Address &x =", &x)
	fmt.Println("value *y =", *y)
	fmt.Println("Address y =", y)

}
