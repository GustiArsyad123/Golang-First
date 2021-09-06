/* File: ConditionUseSwitchCase.go*/
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Switch merupakan seleksi kondisi yang sifatnya fokus pada satu variabel, lalu kemudian
dicek nilainya. Perlu diketahui, switch pada pemrograman Go memiliki perbedaan dibanding
bahasa lain. Pada pemrograman Go, ketika sebuah case terpenuhi, tidak akan dilanjutkan ke
pengecekkan case selanjutnya, meskipun tidak ada kata kunci break. Konsep ini berkebalikan
dengan switch pada umumnya, yaitu ketika sebuah case terpenuhi, maka akan tetap dilanjut
untuk mengecek case selanjutnya kecuali jika ada kata kunci break.
Bentuk umum :
Switch var{
	case nilai1 :
		statement
	case nilai n :
		tatement
	default :
		statement
}
*/
/*---------------*/
package main

import "fmt"

func main() {
	var value = 75
	switch value {
	case 90:
		fmt.Println("Perfect")
	case 75:
		fmt.Println("Good")
	case 60:
		fmt.Println("Not Bad")
	default:
		fmt.Println("almost graduate")
	}
}
