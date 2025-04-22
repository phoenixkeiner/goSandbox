// Decalare a main package
package main

// Import packages
import (
	"fmt" // ability to print
	"math/rand"
	"time" // time
)

func main() {
	// Print Hello, World! to the console
	fmt.Println("Hello, World!")

	// Print the time
	fmt.Println("The time is: ", time.Now())

	// Pick a random number
	fmt.Println("My favorite number is: ", rand.Intn(10))
}
