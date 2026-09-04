package main

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"
)

type Subject struct {
	Name          string
	TotalMarks    float32
	MinMarks      float32
	ObtainedMarks float32
	Grade         string
}

type Format struct {
	Section   string
	Subjects  []Subject
	PassMarks float32
	Year      int
}

type Course struct {
	Id        int
	Name      string
	Syllabus  []Format
	StartYear int
	EndYear   int
}

type Student struct {
	Id         int
	Name       string
	DoB        time.Time
	CourseId   int
	Percentage float32
}

func (st *Student) GetStudents(stud []Student, location *time.Location) []Student {
	stud = append(stud, Student{
		Id:         1,
		Name:       "Girish",
		DoB:        time.Date(1982, time.April, 13, 01, 20, 0, 0, location),
		CourseId:   1,
		Percentage: 78.09,
	})
	return stud
}

type ValidateError struct {
	Field   string
	Message string
}

func (e *ValidateError) Error() string {
	return fmt.Sprintf("Error caused by %s field.\n Message: %s\n", e.Field, e.Message)
}

type Result struct {
	StudentId int
	Err       error
}

type ResultStats struct {
	mu           sync.Mutex
	SuccessCount int
	FailureCount int
}

func (stats *ResultStats) RecordSuccess() {
	stats.mu.Lock()
	defer stats.mu.Unlock()
	stats.SuccessCount++
}

func (stats *ResultStats) RecordFailure() {
	stats.mu.Lock()
	defer stats.mu.Unlock()
	stats.FailureCount++
}

type Reportable interface {
	GenerateReport() string
}

func (rstat *ResultStats) GenerateReport() string {
	reportString := "\t\t\tStudent results\n"
	reportString += fmt.Sprintf("\tTotal Students: %d\n", rstat.SuccessCount+rstat.FailureCount)
	reportString += fmt.Sprintf("\tTotal Passed: %d\n", rstat.SuccessCount)
	reportString += fmt.Sprintf("\tTotal Failed: %d\n", rstat.FailureCount)
	reportString += fmt.Sprintf("\tPass Percentage: %3.2f\n", float64(rstat.SuccessCount/(rstat.SuccessCount+rstat.FailureCount))*100.00)
	return reportString
}

func PrintReport(r Reportable) {
	fmt.Println(r.GenerateReport())
}

func ProcessWorker(ctx context.Context, w int, studChan <-chan Student, resultChan chan<- Result, stats *ResultStats, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case stuDetail, ok := <-studChan:
			if !ok {
				return
			}
			//test for errors
			var err error
			if strings.TrimSpace(stuDetail.Name) == "" {
				err = &ValidateError{Field: "Name", Message: "empty name given"}
			} else if stuDetail.Percentage < 0.0 && stuDetail.Percentage > 100.00 {
				err = &ValidateError{Field: "Percentage", Message: "invalid percentage"}
			}
			if err != nil {
				errMessage := fmt.Errorf("worker %d failed to process student with Id %d\n", w, stuDetail.Id)
				stats.RecordFailure() //TODO: record failure
				resultChan <- Result{StudentId: stuDetail.Id, Err: errMessage}
				continue
			}
			select {
			case <-time.After(200 * time.Millisecond):
				stats.RecordSuccess() //TODO: record success
				resultChan <- Result{StudentId: stuDetail.Id, Err: nil}
			case <-ctx.Done():
				resultChan <- Result{StudentId: stuDetail.Id, Err: ctx.Err()}
				return
			} // end of select if timed out
		} // end of select
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	var location, err = time.LoadLocation("Asia/Kolkata")
	if err != nil {
		panic(err)
	}
	students := []Student{}
	var student *Student
	var results ResultStats
	studChan := make(chan Student, len(students))
	resultChan := make(chan Result, len(students))
	var wg sync.WaitGroup
	// get inputs
	students = student.GetStudents(students, location)
	// start workers
	numWorkers := 2
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go ProcessWorker(ctx, i, studChan, resultChan, &results, &wg)
	} // end of for
	// pass to student channel
	go func() {
		for _, s := range students {
			studChan <- s
		}
		defer close(studChan)
	}()
	// wait for result channel
	go func() {
		wg.Wait()
		defer close(resultChan)
	}()
	// print result
	for rs := range resultChan {
		if rs.Err != nil {
			fmt.Println("Caught Error:", rs.Err)
		} else {
			fmt.Printf("Student Id %d processed successfully\n", rs.StudentId)
		}
	} // end of resultChan
	//generate report
	var reports = []Reportable{&results}
	for _, report := range reports {
		PrintReport(report)
	}
}
