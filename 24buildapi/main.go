package main

import "log"

type Course struct {
	CourserId   string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"`
}

type Author struct {
	AutherName  string
	AutherEmail string
}

var courses []Course

func IsEmpty(c *Course) bool {
	return c.CourseName == ""
}

func main() {
	IsEmpty(&courses)
	console.log('');
	fmt.Println('hey')
	
	
	
	
	
}