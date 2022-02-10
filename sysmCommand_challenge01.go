package main

import (
	"fmt"
	//"io/ioutil"
	"log"
	"os/exec"
	//"strings"
)

func main(){

//	cmd :=  exec.Command ("nslookup google.com")

	//stdout, err := cmd.StdoutPipe()
	
	out, err  :=  exec.Command ("nslookup", "google.com").Output()

	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println (string (out))

/*
	if  err :=  cmd.Start(); err !=nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(stdout)

	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Wait(); err !=nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", string(data))
*/




}


