/* File: SliceFromArray.go*/
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Nilai Dasar Pembentukan Slice
slice dapat dibentuk dari array yang sudah
didefinisikan, dengan teknik 2 indeks untuk mengambil elemen-elemennya.
*/
/*---------------*/
package main

import (
	"fmt"
)

func main() {
	var course = []string{"test", "oke", "i'm", "our", "we", "website", "desktop"}
	var my_course = course[2:6]

	fmt.Println(my_course) //i'm, our, we, website
}
