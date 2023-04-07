package main

import "fmt"

func main() {
	validasi := true
	bilangan := 97

	if bilangan == 0 || bilangan == 1 {
		fmt.Println("Bukan Bilangan Prima")
	} else {
		for i := 2; i < bilangan; i++ {
			if bilangan%i == 0 {
				validasi = false
				break
			}
		}
		if validasi {
			fmt.Println("Bilangan Prima")
		} else {
			fmt.Println("Bukan Bilangan Prima")
		}
	}

}
