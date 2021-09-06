package main

import "fmt"

func main() {
	var first_Name string
	last_Name := "Arsyad"
	fmt.Print("Input your first name:")
	fmt.Scan(&first_Name)
	fmt.Println("Complete your name:" + first_Name + " " + last_Name)
}
