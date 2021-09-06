/* File: Function_strconv_FormatInt.go */
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Fungsi ini digunakan untuk mengkonversi data numerik dengan tipe data int64 ke string dengan
menggunakan basis numerik bisa ditentukan sendiri. */
/*---------------*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num = int64(30)
	var str = strconv.FormatInt(num, 6)
	fmt.Println(str)

}
