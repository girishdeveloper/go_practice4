package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Reportable interface {
	GenerateReport() string
}
type Student struct {
	Id            int
	Name          string
	MarksObtained float32
	Percentage    float32
	CourseId      int
}
type Course struct {
	Id     int
	Name   string
	Format []Subject
}
type Subject struct {
	Id       int
	Name     string
	MaxMarks float32
	MinMarks float32
}

func (s Subject) SetSubjects(totalSubjects int) {

	for i := 1; i <= totalSubjects; i++ {
		subject := Subject{}
		subject.Id = i
		fmt.Print("Enter Subject Name:")
		fmt.Scanf("%s", &subject.Name)
		fmt.Print("Enter Maximum Marks:")
		fmt.Scanf("%f", &subject.MaxMarks)
		fmt.Print("Enter Minimum Marks:")
		fmt.Scanf("%f", &subject.MinMarks)
		subjects = append(subjects, subject)
	} // end of for
}
func (c Course) SetCourse() Course {
	c.Id = 1
	fmt.Print("Enter Course Name:")
	fmt.Scanf("%s", &c.Name)

	c.Format = subjects
	fmt.Printf("Subjects/ format specified for the course are: %v\n", c.Format)
	return c
}
func (c Course) GetCourse(id int) Course {
	return c
}
func (stu Student) Enroll(courseId int, totalStrength int) {
	for i := 1; i <= totalStrength; i++ {
		stud := Student{}
		stud.Id = i
		stud.MarksObtained = 0.00
		stud.Percentage = 0.00
		stud.CourseId = courseId
		fmt.Print(strconv.FormatInt(int64(i), 10), ") Enter student's name:")
		fmt.Scanf("%s", &stud.Name)
		students = append(students, stud)
	}
}
func (c Course) SetScores() {
	for i, student := range students {
		fmt.Println(strings.Repeat("*", 35))
		fmt.Printf("Roll Number: %2d\t\tStudent:%s\n", student.Id, student.Name)
		fmt.Printf("Course: %s\n", c.GetCourse(student.CourseId).Name)
		fmt.Println("Subjects:")
		var totalMaxMarks float32 = 0.00
		for _, subject := range c.Format {
			var marksObtained float32 = 0.00
			fmt.Printf("\n%s:\t%3.2f\t%3.2f\t", subject.Name, subject.MaxMarks, subject.MinMarks)
			fmt.Print("Marks Obtained: ")
			fmt.Scanf("%f", &marksObtained)
			students[i].MarksObtained += marksObtained
			totalMaxMarks += subject.MaxMarks
		} // end of for subjects
		students[i].Percentage = float32(students[i].MarksObtained) / float32(totalMaxMarks) * 100
	} // end of for student
}
func (c Course) GenerateReport() string {
	var reportString = ""
	for i, student := range students {
		if c.Id != student.CourseId {
			continue
		}
		fmt.Println(strings.Repeat("#", 30))
		reportString += "Roll Number: " + strconv.FormatInt(int64(student.Id), 10) + "\t\tStudent:" + student.Name + "\n"
		reportString += "Course: " + c.GetCourse(student.CourseId).Name + "\n"
		reportString += "Subjects:\n"
		for _, subject := range c.Format {
			reportString += subject.Name + ":\t" + strconv.FormatFloat(float64(subject.MaxMarks), 'f', -1, 64) + "\t" + strconv.FormatFloat(float64(subject.MinMarks), 'f', -1, 64) + "\t\n"
		} // end of for subjects
		reportString += "Marks Obtained: " + strconv.FormatFloat(float64(students[i].MarksObtained), 'f', -1, 64) + "\tPercentage: " + strconv.FormatFloat(float64(students[i].Percentage), 'f', -1, 64) + "\n"
	} // end of for student
	return reportString
}
func (stu Student) GenerateReport() string {
	var reportString = ""
	for i, student := range students {
		if stu.Id != student.Id {
			continue
		}
		fmt.Println(strings.Repeat("$", 30))
		reportString += "Roll Number: " + strconv.FormatInt(int64(student.Id), 10) + "\t\tStudent:" + student.Name + "\n"
		reportString += "Course: " + course.GetCourse(student.CourseId).Name + "\n"
		reportString += "Subjects:\n"
		for _, subject := range course.Format {
			reportString += subject.Name + ":\t" + strconv.FormatFloat(float64(subject.MaxMarks), 'f', -1, 64) + "\t" + strconv.FormatFloat(float64(subject.MinMarks), 'f', -1, 64) + "\t\n"
		} // end of for subjects
		reportString += "Marks Obtained: " + strconv.FormatFloat(float64(students[i].MarksObtained), 'f', -1, 64) + "\tPercentage: " + strconv.FormatFloat(float64(students[i].Percentage), 'f', -1, 64) + "\n"
	} // end of for student
	return reportString
}

func PrintReport(r Reportable) {
	fmt.Println(r.GenerateReport())
}

var subjects []Subject
var students []Student
var course Course

func main() {
	var subj Subject
	var student Student
	var totalSubjects, totalStudents int

	fmt.Print("Total Subjects:")
	fmt.Scanf("%d", &totalSubjects)
	fmt.Println("Enter Subject details.")
	subj.SetSubjects(totalSubjects)
	fmt.Printf("Entered subject are: %v\n", subjects)
	fmt.Println("Enter Course details.")
	course = course.SetCourse()
	fmt.Println(course)
ASK_STUDENTS:
	fmt.Print("Total Students to enroll:")
	fmt.Scanf("%d", &totalStudents)
	if totalStudents <= 0 {
		goto ASK_STUDENTS
	}
	fmt.Println("Enroll students in the course.")
	student.Enroll(course.Id, totalStudents)
	course.SetScores()
	reports := []Reportable{course, student}
	for _, report := range reports {
		//fmt.Println(reflect.TypeOf(report), reflect.TypeOf(students[0]))
		if reflect.TypeOf(report) == reflect.TypeOf(students[0]) {
			for _, student := range students {
				PrintReport(student)
			}
		} else {
			PrintReport(report)
		}
	}
}
