package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {

	//1.
	myMap := make(map[string]User)

	first := User{
		FirstName: "Omkar",
		LastName:  "Kulkarni",
	}
	myMap["first"] = first
	
	log.Println(myMap["first"].FirstName)
	log.Println(myMap["second"].LastName)

	// 2nd method:
	myMap2 := map[string]User{
		"first": {
			FirstName: "Omkar",
			LastName:  "Kulkarni",
		},
		"second": {
			FirstName: "John",
			LastName:  "caliber",
		},
	}
	log.Println(myMap2["second"].FirstName)
	log.Println(myMap2["second"].LastName)



	// _________________________________________________________________
	//Slices

	var names[] string
	names = append(names, "Omkar", "John", "chad")
	log.Println(names)

	var nums= []int{5,7,3,1,2}

	log.Println(nums)

	sort.Ints(nums)
	log.Println(nums)


	nums2:= []int{1,2,3,4,5,6,7,8,9,10}

	log.Println(nums2[5:9])




}