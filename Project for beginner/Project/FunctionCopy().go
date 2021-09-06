/* File: FunctionCopy().go*/
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Fungsi copy() digunakan untuk menduplikasi elemen slice pada parameter ke-2 untuk
parameter ke-1, nilai balik dari fungsi copy() adalah jumlah nilai yang berhasil di duplikasi
*/
/*---------------*/
package main

import (
	"fmt"
)

func main() {
	var course = []string{"test", "oke", "i'm", "our", "we", "website", "desktop"}
	var beginner_course = []string{"Cisco", "Linux", "Java", "Golang", "Python", "C Language"}

	var my_course = copy(course, beginner_course)

	fmt.Println(my_course)       //6
	fmt.Println(course)          //[Cisco Linux Java Golang Python C Language desktop]
	fmt.Println(beginner_course) //[Cisco Linux Java Golang Python C Language]

}
