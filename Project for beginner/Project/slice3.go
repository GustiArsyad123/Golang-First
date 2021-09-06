/* File: slice3.go*/
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
jika suatu slice dibentuk dari slice lainnya maka data elemen slice yang baru akan memiliki alamat
memori yang sama dengan slice yang lama. Perubahan pada slice baru tersebut akan
mempengaruhi slice yang lama,
*/
/*---------------*/
package main

import (
	"fmt"
)

func main() {
	var course = []string{"test", "oke", "i'm", "our", "we", "website", "desktop"}
	var my_course = course[2:6]
	fmt.Println(course)
	fmt.Println(my_course)

	my_course[0] = "Databases"
	fmt.Println(course)
	fmt.Println(my_course)
}
