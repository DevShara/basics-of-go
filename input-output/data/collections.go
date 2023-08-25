package data

import "fmt"

//Arrays
var Countries [10]string
var Numbers = [10]int {1,2,3}

//Slices
var Letters = []string {"a", "b", "e"}

//Maps
var Details = map[string]interface{} {"Name": "Ishara", "Age": 24, "City": "Kandy"}

var value int = 1


func init(){
	// Countries[0] = "Sri Lanka"
	// Countries[3] = "India"
	// Numbers[2] = 5
	
	value = 2

	// var qty = cap(Countries)
	fmt.Println("value: ", value)
}



