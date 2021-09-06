/* File: defer.go */
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Defer digunakan untuk mengakhirkan eksekusi sebuah pernyataan, sedangkan Exit digunakan
untuk menghentikan program. Dua bahasan ini digabungkan agar hubungan antara keduanya dapat
dengan lebih mudah dipahami
*/
/*---------------*/
package main

import "fmt"

func main() {
	defer fmt.Println("Done")
	fmt.Println("Good Luck for learn about Golang")
}
