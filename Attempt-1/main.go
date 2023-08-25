// Frist Go program

package main

import "fmt"

func main (){
	// *** First program 
	fmt.Println("Hello World");
	fmt.Println("This is my frist Go program!!!!!!");

	// *** learning escape sequence in Golang
	// escape sequences 
	// single line comment 

	/*
	 multi line comment
	 */

	fmt.Println("This is new line-- \n--escape sequence....")
	fmt.Println("This is tab \t escape sequence....")
	fmt.Println("This is how print froward slash \\ in your program")
	fmt.Println("This is how print quatetion \" in your program \"")

	// ***Variable 
	// var name string
	// var age, salary int32
	// name = "Kamrul"
	// age = 25
	// salary = 100000

	// or 
	// var name = "Kamrul"
	// var age = 25
	// var salary = 100000

	// or 
	name := "Kamrul"
	age := 25
	salary := 100000

	fmt.Println("Name : ", name )
	fmt.Println("Age : ", age )
	fmt.Println("Salary : ", salary )

	// ***Formate Output and constant
	const Name = "Kabir" // can't change this value. Because it is conts value
	_age := 32
	_salary := 2300000

	fmt.Printf("Name: %v\n", Name)
	fmt.Printf("Age: %v\n", _age)
	fmt.Printf("Salary: %v\n", _salary)


	// ***User input
	var firstName, lastName, country string

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your country:")
	fmt.Scan(&country)

	//print
	fmt.Printf("Hello!!\nMy name is %v %v. I am from %v.", firstName, lastName, country)
}