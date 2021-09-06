/* File: variable.go */
/* Author : Gusti Arsyad, email suriv27121994@gmail.com*/
/* Description :
A declared variable of this type is a variable that must be followed or accompanied by certain data types.
This data type will determine a data value stored in the variable. The term data type used in a variable is
called a method "manifest typing".*/
/*---------------*/
package main

import "fmt"

func main() {
	var first_Name, last_Name string
	fmt.Print("Input your first name:")
	fmt.Scan(&first_Name)
	fmt.Print("Input your last name:")
	fmt.Scan(&last_Name)
	fmt.Println("Complete your name:" + first_Name + "" + last_Name)
}
