package main

import "frontendmasters.com/go/server/data"

func main() {
	sam := data.Instructor{}
	sam.FirstName = "Ishara"
	sam.LastName = "Samuditha"

	man := data.NewInstructor("Minura", "Samuditha")
	man.Score = 56

	goCourse := data.Course{Id: 1, Name: "Go basics", Instructor: sam}

	print(sam.Print())
	print(goCourse.String())
}
