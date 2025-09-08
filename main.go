package main

import "fmt"

func displayGreeting(name string, age int) {
	fmt.Println("Hello, cuddly-goggles! My name is", name, "and I am", age, "years old.")

	if name == "david" {
		fmt.Println("Hello, you legend!")
	} else {
		fmt.Println("Nice to meet you!")
	}
}
func performCountdown(start int) {
	fmt.Println("\n--- Starting Countdown ---")
	for i := start; i > 0; i-- {
		fmt.Println("T-minus:", i, "seconds")
	}
	fmt.Println("Liftoff!")
}	
func checkAge(age int) {
	if age < 18 {
		fmt.Println("You are a minor.")
	} else {
		fmt.Println("Age check complete.")
	}
}

func main() {
	myName := "david"
	myAge := 21
	hobbies := []string{"Reading", "Hiking"}
	displayGreeting(myName, myAge)
	checkAge(myAge)
	performCountdown(myAge)
	hobbies = append(hobbies, "Coding", "Gaming")
    fmt.Println("Updated hobbies:", hobbies)
	for index, hobby := range hobbies {
    	fmt.Printf("Hobby #%d is %s\n", index, hobby)
	}
	
	fmt.Println("\n--- End of Program ---") 
}
	
