/* File: for2.go */
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Konsep perulangan ini yaitu menuliskan kondisi setelah kata kunci for (hanya kondisi).
Deklarasi dan iterasi variabel counter tidak dituliskan setelah kata kunci, tetapi hanya kondisi
perulangan saja. Konsepnya mirip seperti while milik bahasa pemrograman lain.
*/
/*---------------*/
package main

import (
	"fmt"
)

func main() {
	var i = 0
	for i < 10 {
		fmt.Println("Number", i)
		i++
	}
}
