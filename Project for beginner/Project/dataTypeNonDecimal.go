/* File: MultiVariable.go */
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Tipe data numerik non-desimal merupakan tipe data yang digunakan untuk nilai data bilangan
bulat. Tipe data numerik non-desimal juga disebut dengan non floating point, di Golang ada dua UINT and INT*/
/*---------------*/
package main

import "fmt"

func main() {
	var positive_numbers uint = 90
	var negativenumbers int = -800

	fmt.Printf("Static positive numbers = %d\n", positive_numbers)
	fmt.Printf("Static negative numbers = %d\n", negativenumbers)
}
