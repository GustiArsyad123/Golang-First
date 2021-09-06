/* File: constanta.go */
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Exit digunakan untuk menghentikan program secara paksa pada saat itu juga. Semua
pernyataan setelah Exit tidak akan dieksekusi, termasuk juga defer.
Fungsi os.Exit() berada dalam package os. Fungsi ini memiliki sebuah parameter bertipe
numerik yang wajib diisi. Angka yang dimasukkan akan muncul sebagai exit status ketika program
berhenti.
*/
/*---------------*/
package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Done")
	os.Exit(1)
	fmt.Println("Good Luck for learn about Golang")

}
