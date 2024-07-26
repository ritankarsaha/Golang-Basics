package main

import (
	f "fmt"
	s "sort"
)

func main(){
   f.Println("Slices in Golang:- ")
   var fruitList = []string{"Apple", "Tomato", "Peach", "Cucumber", "Hello"}
   f.Println("the fruit list there is: ",fruitList)
   fruitList = append(fruitList, "Mango","banana")
   f.Println("The updated fruitlist is :-",fruitList)
   newfruitList := append(fruitList[:3])
   f.Println("The updated fruit list is:- ",newfruitList)
//    highScores := make([]int,2929)
//    highScores[0] = 256
//    highScores[1] = 2561
//    highScores[2] = 2562
//    highScores[3] = 2563
//    highScores[4] = 2564
//    newhighscores := append(highScores, 2232,4322,332,233,2,2,3,2,2)
//    f.Println("The updated highscores is: ",newhighscores)


highScores := []int{2,2,2,2,2,3,3,3,3,2,2,3,2,2,1,3,4,5,6,6,5,5,32,2,1,2,3,34,2,2,1,3,3,2,2,2,3,4,3,3,34,32}
f.Println(s.IntsAreSorted(highScores)) //check whether are sorted
s.Ints(highScores)
f.Println("The sorted array now is:- ",highScores) //sorting the array in golang


lolList := []string{ "lol", "lol1", "lol2", "lol3", "lol4", "lol5", "lol6", "lol7", "lol8"}
f.Println("The entire lol list is:", lolList)
var index int = 2

// Remove the element at the specified index
lolList = append(lolList[:index], lolList[index+1:]...)

f.Println("The lol list after removing element at index", index, "is:", lolList)


}




