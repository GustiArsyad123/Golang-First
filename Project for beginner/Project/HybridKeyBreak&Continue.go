/* File: HybridKeyBreak&Continue.go */
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Kata kunci break digunakan untuk menghentikan secara paksa sebuah perulangan yang
berjalan, sedangkan continue dipakai untuk memaksa program untuk lanjut ke perulangan
berikutnya.
*/
/*---------------*/
package main

import "fmt"

func main() {
	for i := 1; i <= 15; i++ {
		if i%2 == 1 {
			continue
		}
		if i > 14 {
			break
		}
		fmt.Println("Number", i)
	}
}
