package main

import f "fmt"

func main(){
	f.Println("Functions in Golang")
    greeting()

	result := adder(3,5)
	f.Println("The result of the additiong of twio number is: ",result)

	res, msg := multiadder(3,3,3,32,2,4,2,2,3,2,32,3)
	f.Println("The reuslt is ", res)
	f.Println("The message is ",msg)


}

func greeting(){
	f.Println("Hello i am Ritankar Saha")
}
func adder (var1 int ,var2 int) int {
	return var1 + var2
}

func multiadder(values ...int) (int , string){
	total :=0
	for _,val := range values{
		total+=val
	}
	return total, "Hello i am ritankar"
}
