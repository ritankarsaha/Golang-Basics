package main

import f "fmt"


func main(){
 f.Println("Maps in Golang:- ")
 languages := make(map[string]string)
 languages["JS"] = "Javascript"
 languages["TS"] = "typeScript"
 languages["RR"] = "Ruby on Rails"
 languages["GO"] = "Golang"

 f.Println("List of all the languages are: ",languages)
 f.Println("The original for JS is",languages["JS"])
 delete(languages,"RR")
 f.Println("The updated list of languages is ",languages)



}
