/* File: SeleksiKondisiIF-Else.go*/
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Penerapan if-else pada bahasa pemrograman Go, sama seperti pada bahasa
pemrograman lainnya, yang membedakan hanya pada tanda kurungnya (parentheses), di
bahasa pemrograman Go tidak perlu ditulis tanda kurung.
Bentuk umum :
IF kondisi{
	statement
}Else If kondisi{
	statement
}Else If{
	statement
}Else{
	statement
}
*/
/*---------------*/
package main

import "fmt"

func main() {
	var grade = 100
	if grade == 100 {
		fmt.Println("Graduate with perfect value")
	} else if grade > 50 {
		fmt.Println("Graduate")
	} else if grade == 45 {
		fmt.Println("almost graduate")
	} else {
		fmt.Printf("not graduate, your value %d\n", grade)
	}
}
