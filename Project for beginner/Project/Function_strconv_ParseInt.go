/* File: Function_strconv_ParseInt.go */
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Fungsi ini digunakan untuk mengkonversi tipe data string yang berbentuk numerik dengan basis
tertentu ke tipe numerik non-desimal dengan lebar data bisa ditentukan. */
/*---------------*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str = "879"
	var num, err = strconv.ParseInt(str, 10, 64)
	if err == nil {
		fmt.Println(num)
	}

}
