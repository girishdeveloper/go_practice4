package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

type Exam struct {
	subject       string
	marksMaximum  float32
	marksObtained float32
	marksMinimum  float32
	grade         string
}

const EXAM_NAME string = "AISSCE"
const MIN_MARK float32 = 0.00
const MAX_MARK float32 = 100.00
const PASS_MARK float32 = 33.00

type RESULT int

const (
	Fail RESULT = iota //0
	Pass               //1
)

type Result struct {
	rollNumber int
	examName   string
	result     RESULT
	grade      string
}

func main() {
	totalSubjects := 5
	totalResults := 30
	exams := make(chan []Exam, 5)
	results := make(chan Result, 10)
	Subjects := [5]string{"Maths", "Physics", "Chemistry", "Biology", "English"}
	var wg sync.WaitGroup

	for w := 1; w <= totalResults; w++ {
		wg.Add(1)
		go studentWorker(w, exams, results, &wg)
	}
	// record and send
	go func() {
		ex := make([]Exam, totalSubjects)
		for i := 1; i <= totalResults; i++ {
			ex = []Exam{}
			for j := 1; j <= totalSubjects; j++ {
				obtainedMark := float32(MIN_MARK + rand.Float32()*(MAX_MARK-MIN_MARK))
				var grade string = decideGrade(obtainedMark)
				// load data
				ex = append(ex, Exam{
					subject:       Subjects[j-1],
					marksMaximum:  MAX_MARK,
					marksObtained: obtainedMark,
					marksMinimum:  PASS_MARK,
					grade:         grade,
				})
			} // end of for subjects
			exams <- ex
		} // end of for exams
		close(exams)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()
	for k := 1; k <= totalResults; k++ {
		resource := <-results
		finalResult := "pass"
		if resource.result == Fail {
			finalResult = "fail"
		}
		fmt.Printf("Exam result for roll number: %d, %s, result is %s, grade is %s\n",
			resource.rollNumber, resource.examName, finalResult, resource.grade)
	} // end of for totalResults
}

func decideGrade(marksOrPercent float32) (grade string) {
	switch {
	case marksOrPercent >= 90.00:
		grade = "A+"
	case marksOrPercent >= 75.00:
		grade = "A"
	case marksOrPercent >= 60.00:
		grade = "B"
	case marksOrPercent >= 50.00:
		grade = "C"
	case marksOrPercent >= 40.00:
		grade = "D"
	case marksOrPercent >= 33.00:
		grade = "E"
	default:
		grade = "F"
	} // end of switch
	return
}

func calcResult(exam []Exam) (result RESULT, percentage float32) {
	result = Pass
	var totalMarks float32 = 0.00
	var totalMaxMarks float32 = 0.00
	for _, each := range exam {
		totalMarks += each.marksObtained
		totalMaxMarks += each.marksMaximum
		if each.marksObtained < each.marksMinimum {
			result = Fail
			//break
		}
	}
	percentage = (totalMarks / totalMaxMarks) * 100
	return
}

func studentWorker(w int, ex <-chan []Exam, rslt chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	//fmt.Println(w, "Exam given for", len(ex), "subjects")
	for exam := range ex {
		if len(exam) > 0 {
			//fmt.Printf("Result of student %d\n", w)
			time.Sleep(time.Millisecond * 500)
			//fmt.Println("EXAM:", exam)
			rest, percentage := calcResult(exam)
			var overallGrade string = decideGrade(percentage)
			rslt <- Result{
				rollNumber: w,
				examName:   EXAM_NAME,
				result:     rest,
				grade:      overallGrade,
			}
		} else {
			rslt <- Result{
				rollNumber: w,
				examName:   EXAM_NAME,
				result:     0,
				grade:      "F",
			}
		}
	}
}
