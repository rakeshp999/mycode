/* RZFeeser | Alta3 Research
   CHALLEGE 01 - Create a struct named Astro */

package main

import "fmt"

type astro struct {
    name     string
    age      int
    mission  string
    isneeded bool
}

type nasaMission struct{
	people []astro
	number int
	message string
}


func main() {

    ast1 := astro{"Megan McArthur", 35, "ISS", false}
    ast2 := astro{"Barry Wilmore", 41, "Boeing Crew Flight Test (CFT)", true}
    ast3 := astro{"Raja Chari", 39, "SpaceX Crew-3", true}

    fmt.Println(ast1)
    fmt.Println(ast2)
    fmt.Println(ast3)

    	//slice named people made up of astro

	p:= []astro{ast1, ast2,ast3}

	//Display the slice
	fmt.Println(p)

	//select data from the struck

	fmt.Println(p[1].mission)

	// initialize a nasaMission struct, "s"

	s := nasaMission{p, 3,"sucess"}

	//display "s" without fileds

	fmt.Println(s)

	//display "s" with fields
	fmt.Printf("%+v", s)

}
  
