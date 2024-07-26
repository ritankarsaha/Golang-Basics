package main

import f "fmt"

func main(){
	f.Println("Structs in Golnag")
	ritankar:= User{"Ritankar","ritankar.saha786@gmail.com",true,19}
	f.Printf("Details of Ritankar are:- %+v ", ritankar)
	f.Printf("Name is:- %v and email is %v",ritankar.Name,ritankar.Email)

}

type User struct {
	Name string
	Email string
	Status bool
	Age int
}

