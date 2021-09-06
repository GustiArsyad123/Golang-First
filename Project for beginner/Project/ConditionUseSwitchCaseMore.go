/* File: ConditionUseSwitchCaseMore.go*/
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Sebuah case dapat menampung banyak kondisi. Cara penerapannya yaitu dengan
menuliskan nilai pembanding-pembanding variabel yang dilakukan switch setelah kate kunci
case dipisah oleh tanda koma ( , ).
Bentuk umum :
Switch var{
	case nilai1 :
		statement
	case nilai 2,3,4,5,dst :
		statement
	default :
		statement
}
*/
/*---------------*/
package main

import "fmt"

func main() {
	var value = 55
	switch value {
	case 90:
		fmt.Println("Perfect")
	case 70, 75, 80:
		fmt.Println("Good")
	case 55, 60, 65:
		fmt.Println("Not Bad")
	default:
		fmt.Println("almost graduate")
	}
}
