/* File: for.go */
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Perulangan dengan memasukkan variabel counter perulangan.
Cara perulangan ini yaitu memasukkan variabel counter perulangan beserta
kondisinya setelah kata kunci for. Bentuk umum ini hampir sama seperti bahasa
pemrograman C/C++.
*/
/*---------------*/
package main

import (
	"fmt"
)

func main() {
	var i = 0
	for i = 0; i < 10; i++ {
		fmt.Println("Number", i)
	}
}
