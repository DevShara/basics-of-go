package data

import "fmt"

func (i Instructor) Print() string {
	return fmt.Sprintf("%v %v %d", i.FirstName, i.LastName, i.Score)
}

func NewInstructor(firstName string, lastName string) Instructor {
	return Instructor{FirstName: firstName, LastName: lastName}
}

func (c Course) String() string {
	return fmt.Sprintf("%v --- %v --- %v", c.Id, c.Name, c.Instructor)
}
