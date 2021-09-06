/* File: ValueStruct.go*/
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Memberikan nilai diikuti dengan nama dari properti tersebut, cara ini dapat dilakukan
saat pendeklarasian sebuah struct. Urutan properti pada pendeklarasian struct tidak akan
mempengaruhi proses pendeklarasian, contoh:
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
	arsyad := Student{
		Matric_number: "12345",
		Name:          "gusti arsyad",
		age:           25,
	}
	fmt.Printf("Student: %+v\n", arsyad)

}
