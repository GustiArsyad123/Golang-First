/* File: Function_strconv_ParseInt2.go */
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Fungsi ini dapat juga digunakan untuk mengkonversi tipe data dari string “1010” ditentukan
basis numeriknya 2 (biner), akan dionversi ke jenis tipe data int8. */
/*---------------*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str = "1011"
	var num, err = strconv.ParseInt(str, 2, 8)
	if err == nil {
		fmt.Println(num)
	}

}
