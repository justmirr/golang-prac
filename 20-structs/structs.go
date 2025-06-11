package main

import "fmt"

type student struct {
	id    int
	name  string
	marks float64
}

func (s *student) changeMarks(marks float64) {
	s.marks = marks
}

func newStudent(id int, name string, marks float64) *student {
	one := student{
		id:    id,
		name:  name,
		marks: marks,
	}
	return &one
}

func main() {
	two := student{
		id:    111,
		name:  "mir",
		marks: 233.44,
	}

	fmt.Println(two)

	newtwo := 276.55
	two.changeMarks(newtwo)

	fmt.Println(two.marks)

	newone := newStudent(2, "shaukat", 322)

	fmt.Println(newone)

	department := struct {
		name      string
		employees int
	}{"devops", 43}

	fmt.Println(department)
}
