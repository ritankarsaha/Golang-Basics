package main

import (
	f "fmt"
	
)

func main(){

	f.Println("Switch Case statements in golang:- ")
	var number int =6
	switch number {
	case 1:
		f.Println("The number is ",number, "and in statement 1")
	case 2:
		f.Println("The number is ",number, "and in statement 2")
	case 3:
		f.Println("The number is ",number, "and in statement 3")
		fallthrough
	case 4:
		f.Println("The number is ",number, "and in statement 4")
		fallthrough
	case 5:
		f.Println("The number is ",number, "and in statement 5")
	case 6:
		f.Println("The number is ",number, "and in statement 6")
	case 7:
		f.Println("The number is ",number, "and in statement 7")
	case 8:
		f.Println("The number is ",number, "and in statement 8")
	case 9:
		f.Println("The number is ",number, "and in statement 9")
	case 10:
		f.Println("The number is ",number, "and in statement 10")
	}
}

