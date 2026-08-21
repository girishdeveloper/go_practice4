package main

import (
	"fmt"
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

type RESULT int

const (
	Fail RESULT = iota //0
	Pass               //1
)

type Result struct {
	rollNumber int
	examName   string
	result     RESULT
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
				ex = append(ex, Exam{
					subject:       Subjects[j-1],
					marksMaximum:  100,
					marksObtained: 84.50,
					marksMinimum:  33,
					grade:         "B",
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
		fmt.Printf("Exam result for roll number: %d, %s, result is %s\n",
			resource.rollNumber, resource.examName, finalResult)
	} // end of for totalResults
}

func calcResult(exam []Exam) (result RESULT) {
	result = Pass
	for _, each := range exam {
		if each.marksObtained < each.marksMinimum {
			result = Fail
			break
		}
	}
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
			rest := calcResult(exam)
			rslt <- Result{
				rollNumber: w,
				examName:   EXAM_NAME,
				result:     rest,
			}
		} else {
			rslt <- Result{
				rollNumber: w,
				examName:   EXAM_NAME,
				result:     0,
			}
		}
	}
}
