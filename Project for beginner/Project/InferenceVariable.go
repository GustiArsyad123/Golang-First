/* File: InferenceVariable.go */
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Menggunakan kata kunci var dapat dilakukan dalam mendeklarasikan suatu variabel. Bentuk
umum untuk menggunakan deklarsi variabel ini adalah:
var nama_variabel tipe_data
var nama_variabel tipe_data = nilai_data*/
/*---------------*/
package main

import "fmt"

func main() {
	var first_Name string
	last_Name := "Arsyad"
	fmt.Print("Input your first name:")
	fmt.Scan(&first_Name)
	fmt.Println("Complete your name:" + first_Name + " " + last_Name)
}
