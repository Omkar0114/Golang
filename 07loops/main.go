package main

import "log"



type User struct{
	Firstname string
	Lastname string
}

func main() {
	//for loop:

	// for i:=0; i<=1; i++{
	// 	log.Println(i)
	// }
	
	// mySlice := []string{"dog", "horse", "fish ", "banana"}

	// for _, x:= range mySlice{
	// 	log.Println(x)

	// 	//creating a Map:

	// 	myMap := make(map[string]string)

	// 	myMap["dog"] = "dog"
	// 	myMap["fish"] = "fish"
	// 	myMap["hat"] = "hat"

	// 	for _,x= range myMap{
	// 		log.Println(x)
	// 	}

		var mySlice[] User

		u1 := User{
			Firstname: "omkar",

		}

		u2 := User{
			Firstname: "Brad",
		}

		mySlice = append(mySlice, u1)
		mySlice = append(mySlice, u2)


		for i, x := range mySlice{
			log.Println(i,x.Firstname)
		}


	} 
