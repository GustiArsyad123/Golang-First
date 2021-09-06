/* File: Struct.go*/
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Struct pada Golang digunakan untuk menyimpan beberapa properti atau fungsi dalam
1 wadah. Struct sering digunakan dalam pembuatan skema database, skema JSON, skema
XML, dll. Untuk mendeklarasikan sebuah struct dapat menggunakan keyword type diikuti
nama dari struct tersebut dan keyword struct pada Golang,

Pada code di bawah ini, struct Mahasiswa memiliki 3 buah properti yaitu  nomor Matrik  bertipe
string, Nama bertipe string, dan Umur bertipe unsigned integer (bilangan cacah, 0, 1, 2, 3, …,
dst)
*/
/*---------------*/
package main

import "fmt"

type Student struct {
	Matric_number string
	Name          string
	age           uint
}

func main() {
	var arsyad Student
	fmt.Printf("Student: %+v\n", arsyad)

}
