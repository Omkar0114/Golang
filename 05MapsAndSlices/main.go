package main

import "log"

func main() {
	// syntax to create and initialise the map

	// just creating a map
	Mymap := make(map[string]string)

	Mymap["first"] = "Omkar"
	Mymap["second"] = "POP"

	log.Println(Mymap["first"])
	log.Println(Mymap["second"])


	// Declaring and initialising at the same time
	myMap2 := map[string]string{
		"third" : "john",
		"fourth" : "Fido",
	}
	log.Println(myMap2["third"])
	log.Println(myMap2["fourth"])


	myMap3 := map[string]int{
		"Age" : 20,
		"Date" : 30,
	}

	myMap3["Age"] = 21

	log.Println(myMap3["Age"])
	

}