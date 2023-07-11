package main

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

func main() {

}
