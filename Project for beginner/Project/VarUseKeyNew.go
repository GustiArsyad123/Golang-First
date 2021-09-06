/* File: VarUseKeyNew.go */
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Variabel Menggunakan Keyword New*/
/*---------------*/
package main

import "fmt"

func main() {
	name := new(string)
	fmt.Println(name)
	fmt.Println(*name)
}
