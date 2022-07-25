package main

import (
	"fmt"
	"log"
)

// Declaring an interface

type stringer interface{
	String() string
}

// Defining a type

type Article struct{
	Title string
	Author string
}

// creating a reciever function:

func (a Article) String() string{
	return fmt.Sprintf("The %q article was written by %s", a.Title, a.Author)

}

func main(){
	anArticle := Article{
		Title: "Think & Grow rich",
		Author: "Someone",
	}

	log.Println(anArticle.String())
}