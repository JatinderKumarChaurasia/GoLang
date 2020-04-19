package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Student struct {
	name       string
	grade      rune
	percentage float64
	attendDays int
	courses    course
}

type course struct {
	dataLogic    float64
	algorithm    float64
	calculus     float64
	finalProject int
}

const totalDays = 192

type studentList [12]Student

func (l *studentList) generateMarks() {
	for i := 0; i < 12; i++ {
		name := fmt.Sprintf("Student %v", i+1)
		l[i] = Student{
			name:       name,
			grade:      ' ',
			percentage: 0,
			attendDays: 82 + rand.Intn(111),
			courses: course{
				dataLogic:    50 + float64(rand.Intn(51)),
				algorithm:    50 + float64(rand.Intn(51)),
				calculus:     50 + float64(rand.Intn(51)),
				finalProject: 8 + rand.Intn(18),
			},
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var Students studentList
	Students.generateMarks()
	students := Students[:]
	applyBusinessLogic(students)
	sort.Slice(students, func(i, j int) bool {
		return students[i].grade < students[j].grade
	})
	printStudentDetails(students)

}

func printStudentDetails(students []Student) {
	fmt.Printf("\n Student Name \t Percentage \t Attendance \t Project \t Final Grade \t\n")
	fmt.Println("-----------------------------------------------------------------------------")
	for i := 0; i < 12; i++ {
		fmt.Printf(" %-10v \t %8.3v \t %6v \t %12v \t %8c \t\n", students[i].name, students[i].percentage, students[i].attendDays, students[i].courses.finalProject, students[i].grade)
	}
}

func applyBusinessLogic(students []Student) {
	var grade rune
	for i := 0; i < 12; i++ {
		students[i].percentage, grade = calculatePercentageAndGrades(students[i])
		if grade != 'F' {
			calculateAttendance(students[i].attendDays, &grade)
		}
		grade = gradeProject(students[i].courses.finalProject, grade)
		students[i].grade = grade
	}
}

func gradeProject(project int, grade rune) rune {
	percent := float64(project) / 25 * 100
	if percent >= 80 && grade != 'A' {
		grade--
	} else if percent < 40 && grade != 'F' {
		grade++
	}
	return grade
}

func calculatePercentageAndGrades(student Student) (float64, rune) {
	var percentage float64
	grade := 'F'
	percentage = (student.courses.algorithm + student.courses.calculus + student.courses.dataLogic) / 3
	if percentage >= 90.0 {
		grade = 'A'
	} else if percentage < 90.0 && grade >= 80.0 {
		grade = 'B'
	} else if percentage < 80.0 && percentage >= 70.0 {
		grade = 'C'
	} else if percentage >= 60.0 && percentage < 70.0 {
		grade = 'D'
	} else {
		grade = 'F'
	}

	return percentage, grade
}

func calculateAttendance(attendDays int, grade *rune) {
	if float64(attendDays)/totalDays < 0.75 {
		*grade++
	}
}
