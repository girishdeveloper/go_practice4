package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("Validation failed on field %s: %s\n", e.Field, e.Message)
}

type Student struct {
	Id         int
	Name       string
	Percentage float32
}

func (s Student) ValidateStudent() error {
	if strings.TrimSpace(s.Name) == "" {
		return &ValidationError{Field: "Name", Message: "Please enter a valid name."}
	} else if s.Percentage < 0.00 || s.Percentage > 100.00 {
		return &ValidationError{Field: "Percentage", Message: fmt.Sprintf("Invalid percentage %2.2f, percentage should be 0.00 <= x <=100.00.", s.Percentage)}
	}
	return nil
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	student := Student{}
	student.Id = 1
	fmt.Print("Please enter the name:")
	reader.Scan()
	student.Name = reader.Text()
	fmt.Print("Please enter the percentage:")
	fmt.Scanf("%f", &student.Percentage)
	err := student.ValidateStudent()
	if err != nil {
		fmt.Println("Error caught!", err)
		var validErr *ValidationError
		// check the type of error
		if errors.As(err, &validErr) {
			fmt.Printf("--> Inspect error, field %s caused the error.\n", validErr.Field)
		}
		fmt.Println(validErr.Error())
	} else {
		fmt.Printf("%s has scored %2.2f percent.\n", student.Name, student.Percentage)
	}
}
