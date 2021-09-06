/* File: FungsiLen().go*/
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Fungsi len() digunakan untuk menghitung jumlah elemen dari suatu slice
*/
/*---------------*/
package main

import (
	"fmt"
)

func main() {
	var course = []string{"test", "oke", "i'm", "our", "we", "website", "desktop"}

	fmt.Println(len(course)) //7

}
