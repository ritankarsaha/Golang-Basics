package main

import f "fmt"

func main(){
	f.Println("Loops in Golang ")
	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	f.Println(days)
	for i:=0; i<len(days); i++{
		f.Println(days[i])
	}

	for i:= range days{
		f.Println(days[i])
	}

	for _, day := range days {
		f.Println(day)
	}
}

