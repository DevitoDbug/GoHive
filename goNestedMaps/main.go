package main

import (
	"errors"
	"fmt"
)

type StudentGrades map[string]map[string]int

func addNewStudent(grades StudentGrades, name string, subjectGrades map[string]int) (studentGrades map[string]map[string]int) {
	grades[name] = subjectGrades
	return
}

func averageGradeForASubject(subject string, grades StudentGrades) (average float64) {
	numberOfStudents := len(grades)
	sum := 0
	for _, student := range grades {
		_, ok := student[subject] //if the student has grades for the subject
		if !ok {
			continue
		}
		sum += student[subject]
	}

	return float64(sum / numberOfStudents)
}

func averageGradeForStudent(studentName string, grades StudentGrades) (average float64, Error error) {
	student, ok := grades[studentName]
	if !ok {
		return 0.0, errors.New("that student does not exist")
	}

	sum := 0
	numberOfSubjects := len(student)
	for _, subject := range student {
		sum += subject
	}
	return float64(sum / numberOfSubjects), nil
}
func main() {
	classGrades := make(StudentGrades)

	addNewStudent(classGrades, "Odongo", map[string]int{"Math": 100, "English": 90, "Chemistry": 95})
	addNewStudent(classGrades, "Timothy", map[string]int{"Math": 100, "English": 92, "Chemistry": 90})
	addNewStudent(classGrades, "Jane", map[string]int{"Math": 75, "English": 80, "Chemistry": 40})
	addNewStudent(classGrades, "William", map[string]int{"Math": 69, "English": 70, "Chemistry": 60})

	janeAverage, _ := averageGradeForStudent("Jane", classGrades)
	chemistryAverage := averageGradeForASubject("Chemistry", classGrades)

	fmt.Println("Jane's total is: ", janeAverage)
	fmt.Printf("Chemistry results for the class is: %v", chemistryAverage)
}
