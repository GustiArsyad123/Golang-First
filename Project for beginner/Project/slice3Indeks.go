/* File: slice3Indeks.go*/
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
pembentukan slice dapat dilakukan dengan
teknik 3 indeks. Teknik ini digunakan ketika ingin menentukan kapasitas dari slice yang baru
ketika dibentuk, namun nilai dari kapasitas slice yang akan dibentuk tidak boleh melebihi
kapasitas slice yang akan di slicing
*/
/*---------------*/
package main

import (
	"fmt"
)

func main() {
	var course = []string{"test", "oke", "i'm", "our", "we", "website", "desktop"}
	var my_course = course[0:5:5]

	fmt.Println(my_course)      //[test oke i'm our we]
	fmt.Println(cap(my_course)) //5

}
