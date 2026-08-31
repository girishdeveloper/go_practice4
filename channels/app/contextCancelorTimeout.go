package main

import (
	"context"
	"fmt"
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

type StudentResult struct {
	StudentId int
	Data      string
	Err       error
}

type StudentJob struct {
	Ctx      context.Context
	Student  Student
	ResultCh chan<- StudentResult
}

func StudentWorker(ctx context.Context, workerId int, students <-chan Student, result chan<- StudentResult, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case student, ok := <-students:
			if !ok {
				return
			}
			select {
			case <-time.After(500 * time.Millisecond):
				result <- StudentResult{
					StudentId: student.Id,
					Data:      "PASS",
					Err:       nil,
				}
			case <-ctx.Done():
				result <- StudentResult{
					StudentId: student.Id,
					Data:      "TIMEOUT/CANCELLED",
					Err:       ctx.Err(),
				}
				return
			}
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	location, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		panic(err)
	}
	students := []Student{
		{Id: 1, Name: "Girish", DoB: time.Date(1982, time.April, 13, 01, 20, 0, 0, location), CourseId: 1, Percentage: 85.68},
		{Id: 2, Name: "Koresh", DoB: time.Date(1982, time.June, 04, 00, 12, 0, 0, location), CourseId: 1, Percentage: 65.18},
		{Id: 3, Name: "Morish", DoB: time.Date(1982, time.January, 24, 01, 20, 0, 0, location), CourseId: 1, Percentage: 96.88},
	}
	studentChan := make(chan Student, len(students))
	resultChan := make(chan StudentResult, len(students))
	var wg sync.WaitGroup
	totalWorkers := 2
	for w := 1; w <= totalWorkers; w++ {
		wg.Add(1)
		go StudentWorker(ctx, w, studentChan, resultChan, &wg)
	}
	//time.Sleep(10 * time.Second)
	go func() {
		for _, s := range students {
			studentChan <- s
		}
		close(studentChan)
	}()

	go func() {
		wg.Wait()
		close(resultChan)
	}()
	for res := range resultChan {
		if res.Err != nil {
			fmt.Printf("Student %d failed: %v\n", res.StudentId, res.Err)
		} else {
			fmt.Printf("Student %d passed: %s\n", res.StudentId, res.Data)
		}
	}
}
