package main
import "fmt"

// import "frontendmasters.com/go/io/data"

// func calculateTax(price float32)(stateTex float32, cityText float32){
// 	stateTex = price * 0.02 
// 	cityText = price * 0.09
// 	return
// }



func birthday(ageV *int){
	fmt.Printf("The pointer is %v and the value is %v\n", ageV, *ageV)
	*ageV = *ageV + 1
}

func main(){
	age := 22
	birthday(&age)
	birthday(&age)
	birthday(&age)


	// fmt.Print(age)
}
