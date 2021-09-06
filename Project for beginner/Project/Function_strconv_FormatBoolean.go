/* File: Function_strconv_FormatBoolean.go */
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Fungsi ini digunakan untuk mengkonversi tipe data bool ke string. Contoh penerapan dengan
fungsi ini */
/*---------------*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var boolean = true
	var str = strconv.FormatBool(boolean)
	fmt.Println(str)

}
