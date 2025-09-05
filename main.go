package main

import "fmt"

func main() {
	myname := "david"
	myage := 99
    	fmt.Println("Hello, cuddly-goggles! My name is", myname, "and I am", myage, "years old.")
	if myage > 60 {
        	fmt.Println("You are in the golden years!")
    	} else {
        	fmt.Println("Still a spring chicken!")
    	}
	if myname == "david" {
        	fmt.Println("Hello, you legend!")
    	} else {
        	fmt.Println("Nice to meet you!")	
		}	
	for i := 3; i > 0; i-- {
			fmt.Println("T-minus:", i, "seconds")
	}
	fmt.Println("Liftoff!")

}