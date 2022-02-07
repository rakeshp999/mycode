package main

import "fmt"

func  printinfo (name, state string) string{
	return (name + " - " + state)
	//fmt.Println("Print:", name,state)

}
func main(){
	var  fname, state  = "Rakesh" , "NJ"

	 fmt.Println(printinfo(fname,state))
	 fmt.Printf("The type of of name and st is  %T, %T \n", fname, state)

}
