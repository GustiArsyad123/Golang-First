/* File: Function_strconv_ParseBoolean.go */
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Fungsi ini digunakan untuk mengkonversi string ke tipe data bool. Berikut ini contoh penerapan
fungsi strconv.ParseBool */
/*---------------*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str = "true"
	var boolean, err = strconv.ParseBool(str)
	if err == nil {
		fmt.Println(boolean)
	}

}
