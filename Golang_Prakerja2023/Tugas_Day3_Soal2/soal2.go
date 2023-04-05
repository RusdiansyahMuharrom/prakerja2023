package main

import "fmt"

func Mapping(slice []string) map[string]int {
	var hasil = map[string]int{}
	for i := 0; i < len(slice); i++ {
		count := 0
		for j := 0; j < len(slice); j++ {
			if slice[i] == slice[j] {
				count++
			}
		}
		hasil[slice[i]] = count
	}
	return hasil
}

func main() {
	//Tes Case 1
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))

	//Tes Case 2
	fmt.Println(Mapping([]string{}))
}
