/* File: CastingByteStr.go */
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Sebuah [] byte akan dicari bentuk string -nya.
Kode program  merupakan beberapa kode byte string yang dituliskan sebagai
rangkaian slice, yang ditampung oleh variabel b. Kemudian, nilai variabel tersebut dikonversi (cast) ke
dalam tipe data string untuk kemudian ditampilkan dengan perintah fmt.Printf(“%s \n”, s)
*/
/*---------------*/
package main

import (
	"fmt"
)

func main() {
	var b = []byte{87, 101, 108, 99, 111, 109, 101, 32, 116, 111, 32, 71, 111, 108, 97, 110, 103}
	var s = string(b)

	fmt.Printf("%s \n", s)
}
