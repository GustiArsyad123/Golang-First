/* File: FunctionCap().go*/
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Fungsi cap() digunakan untuk menghitung kapasitas maksimum yang tersedia dalam suatu
slice, nilai fungsi cap() akan memiliki nilai seperti fungsi len() ketika pertama dibuat, namun
dapat berubah seiring terjadinya operasi slice.
*/
/*---------------*/
package main

import (
	"fmt"
)

func main() {
	var course = []string{"test", "oke", "i'm", "our", "we", "website", "desktop"}
	fmt.Println(len(course)) //7
	fmt.Println(cap(course)) //7

	var my_course = course[0:4]
	fmt.Println(len(my_course)) //4
	fmt.Println(cap(my_course)) //7

	var arsyad_course = course[2:4]
	fmt.Println(len(arsyad_course)) //2
	fmt.Println(cap(arsyad_course)) //5

}
