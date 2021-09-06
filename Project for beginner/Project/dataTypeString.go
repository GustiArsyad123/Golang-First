/* File: dataTypeString.go*/
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
   Dalam menggunakan tipe data string, nilai data diapit oleh tanda petik dua (“) atau tanda quote.
Selain itu, dapat juga menggunakan tanda grave accent/backticks ( ‘ ), tanda ini terletak di sebelah kiri
tombol 1. Kelebihan string yang dideklarasikan dengan menggunakan grave accent adalah membuat
semua karakter di dalamnya termasuk ke dalam string.*/
/*---------------*/
package main

import "fmt"

func main() {
	var message1 string = "I Love learn Golang "
	var message2 = "I'm Arsyad from Bungin Island"
	fmt.Println(message1, message2)
}
