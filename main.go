package main

import "fmt"


func main() {
	// create the map
	userProfile := map[string]string{
		"name":    "david",
		"favoriteColour": "green",
		"city":    "Imaginaria",
	}		
	fmt.Println("The user's city is:", userProfile["city"])
	
fmt.Println("\n--- End of Program ---") 
}
