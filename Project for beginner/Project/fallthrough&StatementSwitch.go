/* File: fallthrough&StatementSwitch.go*/
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Seperti yang sudah dijelaskan sebelumnya, bahwa konsep switch pada pemrograman
Go memiliki perbedaan dengan bahasa lain. Ketika sebuah case terpenuhi, pengecekkan
kondisi tidak akan diteruskan pada case-case setelahnya. Untuk mengantisipasi agar
pengecekan diteruskan pada case selanjutnya, maka dibutuhkan kata kunci fallthrough untuk
memaksa proses pengecekkan diteruskan pada case selanjutnya.
Bentuk umum :
var nama_var
Switch{
	case nama_var:
		statement
	case (kondisi)
		statement
		fallthrough
	default :
		statement
}
*/
/*---------------*/
package main

import "fmt"

func main() {
	var value = 49
	switch {
	case value == 90:
		fmt.Println("Perfect")
	case (value <= 80) && (value >= 65):
		fmt.Println("Good")
		fallthrough
	case value < 50:
		fmt.Println("You need to learn more")
	default:
		fmt.Println("almost graduate")
		fmt.Println("You can upgrade capacity")
	}
}
