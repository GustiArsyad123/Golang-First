/* File: Slice.go*/
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Slice adalah reference/referensi elemen array, jika suatu array pada dasarnya sudah
memiliki ukuran yang tetap, berbeda dengan slice yang ukurannya dapat secara dinamis dan
mengacu secara fleksibel kepada elemen sebuah array. Slice dapat dibentuk, atau bisa
dihasilkan dari suatu manipulasi sebuah array atau slice lainnya, perubahan pada data di setiap
elemen slice dapat berdampak pada slice lain yang memiliki alamat memori yang sama.
Pembuatan slice memilki cara yang hampir sama seperti pembuatan array, namun
terdapat perbedaan saat melakukan pendefinisian jumlah elemen awal deklarasi.
*/
/*---------------*/
package main

import (
	"fmt"
)

func main() {
	var course = []string{"test", "oke", "saya", "dia", "kamu"}
	fmt.Println(course[0]) //test
}
