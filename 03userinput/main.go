package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	

	welcome := "Welcome to the user input section in golang"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Say somrthing about yourself:-")
	input, _ :=reader.ReadString('\n')
	fmt.Println("You said",input," about yourself")
	fmt.Printf("It's type is %T \n",input)
}
