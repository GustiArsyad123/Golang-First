/* File: CastingStrByte.go */
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Tipe data string sebenarnya adalah slice/array byte. Pada Bahasa pemrograman Golang, sebuah
karakter biasa (bukan Unicode) direpresentasikan oleh sebuah elemen slice byte. Nilai slice tersebut
adalah data int yang (default-nya) berbasis desimal, yang merupakan kode ASCII dari karakter biasa
tersebut.
Cara mendapatkan slice byte dari sebuah data string adalah dengan menggunakan Casting ke
tipe [ ] byte. Tiap elemen byte memiliki isi data numerik dengan basis desimal. Berikut ini adalah
contoh string dalam variabel teks1 yang dikonversi ke dalam [ ] byte. Tiap elemen slice byte tersebut
kemudian ditampilkan satu per satu */
/*---------------*/
package main

import (
	"fmt"
)

func main() {
	var teks_1 = "Welcome to Golang"
	var a = []byte(teks_1)

	fmt.Printf("%d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d\n",
		a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11], a[12], a[13], a[14], a[15], a[16])
}
