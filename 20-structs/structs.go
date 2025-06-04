package main

import "fmt"

type student struct {
	id    int
	name  string
	marks float64
}

func main() {
	one := student{
		id:    111,
		name:  "Mir",
		marks: 214.22,
	}
	fmt.Println(one)

	two := student{
		id:    7,
		name:  "Bond",
		marks: 239.4,
	}
	fmt.Println(two)
	fmt.Println(one.name, two.name)
}
