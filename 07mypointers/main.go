package main

import (f "fmt"
"bufio"
"os"
s "strings"
sc "strconv")

func main(){

	//taking input of the number.
	f.Println("Pointers in Golang:- ")
	f.Println("Please enter a number")
	reader:= bufio.NewReader(os.Stdin)
	input, _:= reader.ReadString('\n')
	input = s.TrimSpace(input)
	number,err := sc.Atoi(input)
	if(err == nil){
		f.Println("The number is: ",number)

	}else{
		f.Println("The error is: ",err)
	}
   
	//pointers in golang
	var ptr = &number
	f.Println("the value of the pointer is: ",ptr)
	f.Println("the actual value of the pointer is: ",*ptr)
	*ptr = *ptr*2
	f.Println("the actual value of the pointer now is: ",*ptr)

}