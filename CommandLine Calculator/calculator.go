package main

import ("fmt"
		// "strconv"
	)

func main(){
	
	calc()
}

func calc () {
	var num1, num2 float64
	var operator string

	fmt.Println("Enter 1st number : ")
	fmt.Scanln(&num1)

	fmt.Println("Enter 2nd number : ")
	fmt.Scanln(&num2)

	fmt.Println("Enter the operator you want : ")
	fmt.Scanln(&operator)


	var ans float64
	switch operator{
	case "+":
		ans = num1 + num2
		fmt.Println("Your answer is ", ans)
	case "-":
		ans = num1 - num2
		fmt.Println("Your answer is ", ans)
	case "*":
		ans = num1 * num2
		fmt.Println("Your answer is ", ans)
	case "/":
		if (num2 == 0){
			fmt.Println("Can not divide by 0!")
		} else {
		ans = num1 / num2
		fmt.Println("Your answer is ", ans)
		}
	default:
		fmt.Println("Invalid operator!, Please try again!")
	}

}