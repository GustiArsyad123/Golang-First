/* File: DataConUseCasting.go */
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Konversi data dengan menggunakan metode Casting adalah konversi tipe data dengan
menggunakan keyword atau kata kunci. Penggunaan konversi ini dengan memanggil tipe data sebagai
fungsi dan menyisipkan nilai data yang akan dikonversi sebagai parameter. */
/*---------------*/
package main

import (
	"fmt"
)

func main() {
	var number1 float64 = float64(25.5)
	fmt.Println(number1)

	var number2 int32 = int32(15.00)
	fmt.Println(number2)
}
