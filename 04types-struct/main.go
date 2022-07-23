package main

import (
	"fmt"
	"log"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	Age       int
	PhoneNo   string
	Birthdate time.Time
}

type Marks struct{
	English string
	DBMS string
	CN int
	AI int
}

func main() {
	User:= User {
		FirstName: "Omkar",
		LastName: "Kulkarni",
		Age: 20,
		PhoneNo: "112-121-212",
	}

	log.Println(User.FirstName, User.LastName)
	log.Println(User.Age)
	log.Println(User.Birthdate)
	fmt.Println(User.PhoneNo)


	Marks:=Marks{
		English: "100",
		CN: 90,
		DBMS: "91",
		AI: 99,
	}

	log.Println(Marks.English, Marks.CN, Marks.DBMS, Marks.AI)
}
