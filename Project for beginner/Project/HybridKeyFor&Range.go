/* File: HybridKeyFor&Range.go */
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Cara ini digunakan untuk melakukan perulangan data bertipe array yang jumlah
itemnya sudah ditentukan.
*/
/*---------------*/
package main

import (
	"fmt"
)

func main() {
	var animals = [5]string{"leon", "Bird", "Corocodile", "Goat", "Chicken"}
	for i, animal := range animals {
		fmt.Printf("item %d : %s\n", i, animal)
	}
}
