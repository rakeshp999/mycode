/* Alta3 Research | RZFeeser
   Variables - Package and function level   */

package main

import "fmt"

var palm, python, tiger bool 

var rock, scissors int = 42, 55

func main() {
	var i int
	var sun, moon = "esclipse", "waxing"
	
	fmt.Println(i, palm, python, tiger)

	fmt.Println("The vlaue of the var rock is:", rock)
	fmt.Println("The value of the var scissors is :", scissors)
	 
	fmt.Println("Look at the moon when it is ", moon + ".")
	fmt.Println("A totla", sun, "of the Sol.")
}

