package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println("Date and time program.")
	fmt.Println(strings.Repeat("#", 13))
	location, _ := time.LoadLocation("Asia/Kolkata")
	var start = time.Now()
	//time.Sleep(time.Second * 15)
	wokeUp := time.Now()
	futureTime := time.Now().Add((10 * time.Second))
	fmt.Println("Started at", start.Format(time.DateTime))
	fmt.Println("Woke up at", wokeUp.Format(time.DateTime))
	fmt.Println("time since start is", time.Since(start))
	fmt.Println("time since awakening is", time.Since(wokeUp))
	fmt.Println("time elapsed since start is", time.Now().Sub(start))
	fmt.Println("time difference is", wokeUp.Sub(start))
	fmt.Println("time until is", time.Until(futureTime))
	fmt.Println("timezone is", location.String())
	fmt.Println(strings.Repeat("#", 13))
	year, month, day := time.Now().Date()
	fmt.Printf("Current Year is:%d\n", year)
	fmt.Printf("Current Month is:%d\n", month)
	fmt.Printf("Current Day is:%d\n", day)
	fmt.Printf("January month value is: %d\n", time.January)
	fmt.Println("making date")
	anyDate := time.Date(2026, time.Month(month), 8, 11, 18, 49, 0, location)
	fmt.Printf("Prepare a date: %v; %v\n", anyDate.Format(time.DateTime), anyDate.Format(time.UnixDate))
	year1, month1, day1 := anyDate.Date()
	fmt.Printf("extract date elements: %d %d, %d\n", year1, month1, day1)
	fmt.Printf("Epoch time is %d\n", time.Now().Unix())
	// dates
	fmt.Println(time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()+1, 0, 0, 0, 0, location))

	fmt.Println(time.Now().AddDate(time.Now().Year(), int(time.Now().Month())+3, time.Now().Day()+3))
	fmt.Println(time.Date(time.Now().Year(), time.Now().Month()+3, time.Now().Day()+3, 0, 0, 0, 0, location))
}
