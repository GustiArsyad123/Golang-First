/* File: dataTypeNumericDecimal.go */
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Tipe data numerik decimal digunakan untuk bilangan yang memiliki angka dibelang koma, atau
bilangan pecahan. Tipe data jenis ini ada dua, yaitu float32 dan float64. Perbedaan kedua tipe data ini
terletak pada jangkauan nilai. Jangakauan nilai ini merujuk ke spesifikasi yang telah diatur dan
dikelauarkan oleh IEEE-754 32-bit floating-point numbers.*/
/*---------------*/
package main

import "fmt"

func main() {
	var bilanganPecahanDecimal = 8.8

	fmt.Printf("Decimal Fractionle Numbers = %f\n", bilanganPecahanDecimal)
	fmt.Printf("Decimal Fractionle Numbers = %.3f\n", bilanganPecahanDecimal)

}
