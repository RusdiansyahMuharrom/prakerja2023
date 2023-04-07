package main

import "fmt"

func main() {
	//Penggunaan scanner
	var a int
	fmt.Print("Masukan nilai a = ")
	fmt.Scanln(&a)
	b := 7
	tinggi := 8

	luas := float64(((a + b) * tinggi) / 2)

	fmt.Println("Luas = ", luas)
}
