/* File: Function_strconv_ParseFloat.go */
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Fungsi dapat digunakan untuk mengkonversi string ke numerik decimal dengan lebar data yang
bisa ditentukan. */
/*---------------*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str = "87.98"
	var num, err = strconv.ParseFloat(str, 32)
	if err == nil {
		fmt.Println(num)
	}

}
