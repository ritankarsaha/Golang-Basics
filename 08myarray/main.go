package main

import f "fmt"



func main(){
	f.Println("Arrays in Golang:- ")
	var veglist = [5]string{"potato","carrot","cucumber","mushroom"}
	f.Println("Vegetable list is ",veglist)
	f.Println("The length of the vegetable list is ",len(veglist))
}
