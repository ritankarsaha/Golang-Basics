package main

import "fmt"

const LoginToken string = "ghabbhhjd" // Public

func main() {
 
	var username = "Ritankar"
	fmt.Println("My name is: ",username)
	fmt.Printf("Variable is of the type %T \n", username)

	var isLoggedIn bool = false
	fmt.Println("The value of the variable is: ",isLoggedIn)
	fmt.Printf("The type of the variable is: %T \n",isLoggedIn)

    var smallval uint8 = 225
	var smallfloat float64 = 225.829829292992
	fmt.Println("The value of the small val is: ",smallval)
	fmt.Println("The value of the small float is: ",smallfloat)
	fmt.Printf("The type is %T \n",smallfloat)
	fmt.Printf("The type is: %T \n",smallval)


	var anotherVariable int = 32
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	var website = "learncodeonline.in"
	fmt.Println(website)

	//usage of the walrus operator in  Golang

	numberOfUser := 300000.0
    username1 := "Ritankar"

    fmt.Println(numberOfUser)
    fmt.Println(username1)

	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
