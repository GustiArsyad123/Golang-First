/* File: forNonArgument&Condition.go */
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Cara berikut ini adalah perulangan dengan kata kunci for ditulis tanpa kondisi. Dengan
ini akan dihasilkan perulangan tanpa henti (sama dengan for bernilai true ). Pemberhentian
perulangan dapat dilakukan dengan menggunakan bantuan kata kunci break
*/
/*---------------*/
package main

import (
	"fmt"
)

func main() {
	var i = 0
	for {
		fmt.Println("Number", i)
		i++
		if i == 10 {
			break
		}
	}
}
