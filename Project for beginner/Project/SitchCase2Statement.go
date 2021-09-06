/* File: SitchCase2Statement.go*/
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Tanda kurung kurawal ( { } ) dapat diterapkan pada kata kunci case dan default . Tanda
kurung kurawal ini bersifat opsional, dapat digunakan juga dapat tidak digunakan. Baiknya jika
dipakai pada blok kondisi yang didalamnya terdapat banyak statement, maka kode akan terlihat
lebih rapih dan mudah dilakukan perawatan.
Bentuk umumnya sama seperti statement switch case biasa, yang berbeda adalah
penggunaan dari statement default -nya saja
*/
/*---------------*/
package main

import "fmt"

func main() {
	var value = 58
	switch value {
	case 90:
		fmt.Println("Perfect")
	case 70, 75, 80:
		fmt.Println("Good")
	case 55, 60, 65:
		fmt.Println("Not Bad")
	default:
		fmt.Println("almost graduate")
		fmt.Println("You can upgrade capacity")
	}
}
