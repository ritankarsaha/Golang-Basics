package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)



func main(){
 fmt.Println("Please give a rating to yourself")
 reader := bufio.NewReader(os.Stdin)
 input, _:=reader.ReadString('\n')
 fmt.Println("The rating that you gave yourself: ",input)
 numRating, err := strconv.ParseFloat(strings.TrimSpace(input),64)
	if err != nil {
		fmt.Println("The error is:- ",err)

	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}
}
