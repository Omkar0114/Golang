package main

import "log"

func main() {

	var myString string = "Omkar"
	log.Println("myString is set to", myString)

	// Passing the reference to 'myString' variable, to access its value
	ChangeUsingPointer(&myString)

	log.Println("After change,myString ia set to", myString)
}

func ChangeUsingPointer(str *string){ //str is a pointer variable of type string
	log.Println("str is set to", str)
	newValue := "Kulkarni"
	*str = newValue
}