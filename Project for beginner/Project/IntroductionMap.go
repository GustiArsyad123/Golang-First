/* File: IntroductionMap.go*/
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Map adalah tipe data asosiatif yang ada di Golang, berbentuk key-value . Untuk setiap
data (atau value ) yang disimpan, disiapkan juga key-nya. Key memiliki sifat unik, karena
digunakan sebagai penanda (atau identifier) untuk pengaksesan value yang bersangkutan.
Kalau dilihat, map mirip seperti slice, hanya saja indeks yang digunakan untuk pengaksesan
bisa ditentukan sendiri tipe-nya (indeks tersebut adalah key)
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
