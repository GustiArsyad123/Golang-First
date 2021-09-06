/* File: SwitchStyleIf-Else.go*/
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Salah satu hal unik pada pemrograman Go yaitu, statement switch dapat digunakan
dengan gaya seperti statement if-else. Nilai yang akan dibandingkan tidak dituliskan setelah
kata kunci switch, melainkan akan ditulis langsung dalam bentuk perbandingan dalam kata
kunci case.
Bentuk umum :
var nama_var
Switch{
	case nama_var:
		statement
	case (kondisi)
		statement
	default :
		statement
}
*/
/*---------------*/
package main

import "fmt"

func main() {
	var value = 50
	switch {
	case value == 90:
		fmt.Println("Perfect")
	case (value <= 80) && (value >= 65):
		fmt.Println("Good")
	default:
		fmt.Println("almost graduate")
		fmt.Println("You can upgrade capacity")
	}
}
