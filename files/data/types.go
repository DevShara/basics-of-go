package data

import "fmt"

type distance float64

// type distanceKm float64

// func (miles distance) toKm() distanceKm {
// 	return distanceKm(1.60934 * miles)
// }

// func (kms distanceKm) toMiles() distance {
// 	return distance(kms / 1.60934)
// }

// func Test() {
// 	d := distance(54.5)
// 	km := d.toKm()
// 	miles := km.toMiles()

// 	fmt.Println("km:", km)
// 	fmt.Println("miles:", miles)
// }

type location string

func (origin location) distanceTo(dest location) {
	fmt.Printf("Origin %v Destination %v", origin, dest)
}

func Test() {
	nyc := location("234")
	nyc.distanceTo("87")
}
