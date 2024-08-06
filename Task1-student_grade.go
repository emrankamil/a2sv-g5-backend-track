package main

import (
	"fmt"
)

func Average(grades map[string]float64) float64{
	var total float64 = 0
	for _, grade := range grades{
		total += grade
	}

	return total / float64(len(grades))
}

func GradeCalculator() {
	var name string 
	var number_of_sub int64
	var grades = make(map[string]float64)

	fmt.Print("name: ")
	fmt.Scanln(&name)
	fmt.Print("Number of subjects: ")
	fmt.Scanln(&number_of_sub)

	for i:= range(number_of_sub){
		fmt.Printf("input name and grade of subject %v \n", i + 1)
		var subject string
		var grade float64
		fmt.Print("subject: ")
		fmt.Scanln(&subject)
		for range(3){
			fmt.Print("grade: ")
			fmt.Scanln(&grade)
			if grade < 0 || grade > 100{
				fmt.Println("grade should be in the range 0 - 100")
			}else{
				break
			}
		}

		grades[subject] = grade
	}

	averageGrade := Average(grades)

	fmt.Printf("Student Name: %s\n\n", name)
	fmt.Println("Subject Grades:")
	for subject, grade := range grades {
		fmt.Printf("- %s: %.2f\n", subject, grade)
	}
	fmt.Printf("\nAverage Grade: %.2f\n", averageGrade)

}