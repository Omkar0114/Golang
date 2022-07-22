package main

import (
	"fmt"
)

func main() {

	// Declaring variable
	var WhatToSay string = "Goodbye"

	//printing the String
	fmt.Println(WhatToSay)

	//Declaring integer
	var i int = 10

	//printing integer
	fmt.Println("Number is set to =" , i)


	//calling a function:
	WhatWasSaid , otherthing := something()
	fmt.Println(WhatWasSaid , otherthing)
	
}

func something() (string , string){
	return "Omkar", "2ndString"
}