/* File: Function_strconv_FormatFloat.go */
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Fungsi ini digunakan untuk mengkonversi data bertipe float64 ke string dengan format eksponen,
lebar digit desimal, dan lebar tipe data bisa ditentukan. */
/*---------------*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num = float64(87.98)
	var str = strconv.FormatFloat(num, 'f', 6, 64)

	fmt.Println(str)

}
