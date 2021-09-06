/* File: SliceWithMake.go*/
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
suatu slice dapat juga dibentuk dengan menggunakan keyword make,
pembuatan slice dengan keyword make ini akan membuat suatu array yang dinamis. Pada
contoh dibawah mungkin tidak akan terlihat perbedaanya dengan array pada umumnya, namun
akan terlihat ketika kode yang akan digunakan lebih kompleks.
*/
/*---------------*/
package main

import (
	"fmt"
)

func main() {
	var urang_course = make([]string, 10)

	fmt.Println(urang_course) // [  ]

}
