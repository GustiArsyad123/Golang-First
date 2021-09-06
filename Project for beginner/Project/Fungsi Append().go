/* File: FunctionAppend().go*/
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Fungsi append() digunakan untuk menambahkan elemen pada slice. Elemen yang baru
ditambahkan akan di posisikan setelah indeks paling akhir. Dan nilai balik dari append() adalah
slice yang sudah ditambahkan nilai baru
*/
/*---------------*/
package main

import (
	"fmt"
)

func main() {
	var course = []string{"test", "oke", "i'm", "our", "we", "website", "desktop"}

	var my_course = course[1:5]

	my_course = append(my_course, "Cryptography")
	fmt.Println(len(my_course)) //5
	fmt.Println(cap(my_course)) //6

	fmt.Println(my_course) //[oke i'm our we Cryptography]
	fmt.Println(my_course) //[oke i'm our we Cryptography]

}
