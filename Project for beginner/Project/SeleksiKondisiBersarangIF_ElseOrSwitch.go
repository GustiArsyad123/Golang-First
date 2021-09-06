/* File: SeleksiKondisiBersarangIF_ElseOrSwitch.go*/
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Seleksi kondisi bersarang merupakan seleksi kondisi yang berada didalam seleksi
kondisi lain, yang mungkin juga berada dalam seleksi kondisi lain, dan seterusnya. Seleksi
kondisi bersarang bisa dilakukan pada statement if - else, switch, ataupun kombinasi keduanya,
Bentuk umum :
var nama_var
If kondisi{
	Switch nama_var{
		case nama_var:
			statement
		default :
		statement
		}
} Else {
	If kondisi{
		statement
} Else If kondisi{
		statement
} Else {
		statement
	If kondisi{
		statement
		}
	}
}
*/
/*---------------*/
package main

import "fmt"

func main() {
	var value = 100
	if value > 75 {
		switch value {
		case 100:
			fmt.Println("Perfect")
		default:
			fmt.Println("nice")
		}
	} else {
		if value == 50 {
			fmt.Println("not bad")
		} else if value == 40 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("You could do it")
			if value == 0 {
				fmt.Println("Try Harder")
			}
		}
	}
}
