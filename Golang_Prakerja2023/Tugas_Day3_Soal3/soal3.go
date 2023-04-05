package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {
	var hasil []int
	for _, c := range angka {
		count := 0
		for _, c2 := range angka {
			if c == c2 {
				count++
			}
		}
		if count == 1 {
			temp, _ := strconv.Atoi(string(c))
			hasil = append(hasil, temp)
		}
	}

	return hasil
}

func main() {
	fmt.Println(munculSekali("1234321"))
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("1122334455"))
	fmt.Println(munculSekali("0872504"))
}
