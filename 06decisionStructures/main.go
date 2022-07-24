package main

import "log"

func main() {
	//if-else statement

	isTrue := true

	if isTrue {
		log.Println("isTrue is set to" ,  isTrue)
	} else {
		log.Println("isTrue is set to" , isTrue)
	}


	// switch statements:

	myAge := 20

	switch myAge{
	case 18:
		log.Println("You cannot vote, Age:=", myAge)

	case 19:
		log.Println("You cannot vote, Age:=", myAge)

	case 20:
		log.Println("You can vote, Age:=", myAge)

	default:
		log.Println("Invalid input!")
	}
}