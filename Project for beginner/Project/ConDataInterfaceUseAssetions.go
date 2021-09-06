/* File: ConDataInterfaceUseAssetions.go */
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
Konversi data dengan teknik type assertions merupakan teknik casting data interface{} ke
segala jenis tipe (dengan syarat datatersebut memang bisa di-casting).
Berikut merupakan contoh penerapannya. Disiapkan variabel nilai_data bertipe
map[string]interface{} dengan value berbeda-beda tipe datanya.*/
/*---------------*/
package main

import (
	"fmt"
)

func main() {
	var dataValue = map[string]interface{}{
		"name":         "gusti arsyad",
		"grade":        90,
		"weigth":       55,
		"heigth":       165.5,
		"genderIsmale": true,
		"myHobbies":    []string{"travelling", "study", "and", "diving"},
	}

	fmt.Println(dataValue["name"].(string))
	fmt.Println(dataValue["grade"].(int))
	fmt.Println(dataValue["weigth"].(int))
	fmt.Println(dataValue["heigth"].(float64))
	fmt.Println(dataValue["genderIsmale"].(bool))
	fmt.Println(dataValue["myHobbies"].([]string))
}
