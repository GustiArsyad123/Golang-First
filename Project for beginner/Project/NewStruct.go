/* File: NewStruct.go*/
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Menggunakan fungsi new, cara ini akan mengalokasikan memori dan mengembalikan
nilai bertipe pointer dengan nilai default pada setiap propertinya sama dengan kosong
(0 untuk int, 0.0 untuk float, “” untuk string, nil untuk pointer, dst)
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
	arsyad := new(Student)
	fmt.Printf("Student: %+v\n", arsyad)

}
