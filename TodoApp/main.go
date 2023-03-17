package main

import "fmt"

func main(){

	todos := make([] string, 0)

	for{
		var choice string
		fmt.Println("1. Add Todo")
		fmt.Println("2. View Todo")
		fmt.Println("3. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)


		switch choice{
		case "1":
			var todo string
			fmt.Print("Enter todo:")
			fmt.Scanln(&todo)
			todos = append(todos, todo)
			fmt.Println("Todo added successfully")

		case "2":
			fmt.Println("Your todos are:")
			for i, todo := range todos{
				fmt.Println(i+1, todo)
			}
		case "3":
			fmt.Println("Thank you for using our app")
			return

		default:
			fmt.Println("Invalid choice")

		}

	}
}