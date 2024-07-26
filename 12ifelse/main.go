package main

import f "fmt"

func main(){
	f.Println("if else statements in golang")
	ritankar := 10
	if ritankar >= 10{
      f.Println("user greater than 10")
	}else{
     f.Println("Number not greater than 10")
	}
	

	if 9%2 == 0{
		f.Println("Number is even")
	}else{
		f.Println("Number is odd")
	}
}
