package main

import (
	"fmt"
)

func getMinMax(numbers ...*int) (min int, max int) {
	var temp_min int = *numbers[0]

	var temp_max int = *numbers[0]

	for i := 0; i < 5; i++ {
		if temp_max < *numbers[i] {
			temp_max = *numbers[i]
		}

		if temp_min > *numbers[i] {
			temp_min = *numbers[i]
		}
	}

	return temp_min, temp_max
}

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Scanln(&a1)
	fmt.Scanln(&a2)
	fmt.Scanln(&a3)
	fmt.Scanln(&a4)
	fmt.Scanln(&a5)
	fmt.Scanln(&a6)

	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)

	fmt.Println("Nilai min ", min)
	fmt.Println("Nilai max", max)

}
