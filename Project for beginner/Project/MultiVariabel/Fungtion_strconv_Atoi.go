/* File: Function_strconv_Atoi.go */
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Fungsi ini dapat digunakan untuk mengkonversi data dari tipe string ke int. Mengembalikan 2
buah nilai balik, yaitu hasil konversi dan error (jika konversi sukses, maka error akan berisi nil). */
/*---------------*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str = "879"
	var num, err = strconv.Atoi(str)
	if err == nil {
		fmt.Println(num)
	}
}
