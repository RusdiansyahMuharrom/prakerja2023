package main

import "fmt"

type Student struct {
	name  string
	score int
}

func main() {
	fmt.Println("Input : ")

	var students []Student
	var total int = 0

	var temp_min int = 999
	var idx_min int = 0

	var temp_max int = 0
	var idx_max int = 0

	for i := 0; i < 5; i++ {
		students = append(students, Student{"", 0})

		fmt.Print("Input ", i+1, " Student's Name ")
		fmt.Scanln(&students[i].name)

		fmt.Print("Input ", i+1, " Student's Score ")
		fmt.Scanln(&students[i].score)

		if temp_max < students[i].score {
			temp_max = students[i].score
			idx_max = i
		}

		if temp_min > students[i].score {
			temp_min = students[i].score
			idx_min = i
		}

		total += students[i].score
	}

	fmt.Println("Output :")
	fmt.Println("Average Score :", float64(total)/float64(len(students)))
	fmt.Printf("Min score of Students : %s (%d) \n", students[idx_min].name, temp_min)
	fmt.Printf("Max score of Students : %s (%d) \n", students[idx_max].name, temp_max)
}
