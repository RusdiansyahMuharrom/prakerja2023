package main

import "fmt"

func main() {
	bilangan := 77

	if bilangan%7 != 0 {
		fmt.Println("Bukan Bilangan Kelipatan 7")
	} else {
		fmt.Println("Bilangan Kelipatan 7")
	}
}
