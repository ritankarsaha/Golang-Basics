package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	username := User{"Ritankar", "ritankar.saha786@gmail.com", true, 16}
	fmt.Println(username)
	fmt.Printf("hitesh details are: %+v\n", username)
	fmt.Printf("Name is %v and email is %v.\n", username.Name, username.Email)
	username.GetStatus()
	username.NewMail()
	fmt.Printf("Name is %v and email is %v.\n", username.Name, username.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "ritankar.saha786@gmail.com"
	fmt.Println("Email of this user is: ", u.Email)
}
