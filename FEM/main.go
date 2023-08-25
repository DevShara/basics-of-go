package main

import (
	"frontendmasters.com/go/server/data"
)

func main() {
	sam := data.Instructor{}
	sam.FirstName = "Ishara"
	sam.LastName = "Samuditha"
	sam.Score = 56

	// man := data.NewInstructor("Minura", "Samuditha")

	goCourse := data.Course{Id: 1, Name: "Go basics", Instructor: sam}

	// print(sam.Print())
	print(goCourse.String())

	// swiftWS := data.NewWorkshop("Swift for iOS", sam)

	// print(swiftWS.String())
}
