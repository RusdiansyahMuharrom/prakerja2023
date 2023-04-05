package main

import (
	"fmt"
)

func ArrayMerge(arrayA, arrayB []string) []string {
	arrayA = append(arrayA, arrayB...)
	var temp = map[string]int{}
	for i, A := range arrayA {
		temp[A] = i
	}

	var hasil []string
	for key := range temp {
		hasil = append(hasil, key)
	}
	return hasil
}

func main() {
	//Test case 1
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))

	//Test case 2
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "brian"}))

	//Test case 3
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))

	//Test case 4
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))

	//Test case 5
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))

	//Test case 6
	fmt.Println(ArrayMerge([]string{}, []string{}))
}
