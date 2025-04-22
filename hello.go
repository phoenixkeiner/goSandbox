// Decalare a main package
package main

// Import packages
import (
	"fmt"       // ability to print
	"math"      // math...
	"math/rand" // random number
	"time"      // time
)

func main() {
	// Print Hello, World! to the console
	fmt.Println("Hello, World!")

	// Print the time
	fmt.Println("The time is: ", time.Now())

	// Pick a random number
	fmt.Println("My favorite number is: ", rand.Intn(10))

	// %g format verb is used to formate a floating point number
	fmt.Printf("Now you have %g problems. \n", math.Sqrt(7)) // calculates the sqaure root of 7

	// Println is built in function, no formatting options, auto new line, mainly for debugging
	// Printf comes from fmt package. Does not automatically add a newline unless you include \n

	fmt.Printf("Hello %s, you are %d years old.\n", "Phoenix", 28)

	name := "Phoenix"
	age := 25
	score := 92.5

	fmt.Printf("Name: %s, Age: %d, Score: %.1f%%\n", name, age, score)
}
