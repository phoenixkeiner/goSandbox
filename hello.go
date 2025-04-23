// Decalare a main package
package main

// Import packages
import (
	"fmt"  // ability to print
	"math" // math...
	"math/cmplx"
	"math/rand" // random number
	"time"      // time
)

// function to add 2 numbers
// can also do add(x, y int) when two or more
// consecutive named function parameters share a type
func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	// A return statement without arguments returns
	// the named return values. This is known as a
	// "naked" return.
	return
}

// The var statement declares a list of variables; as in function argument lists,
// the type is last.
var n, j int = 1, 2
var c, python, java bool

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

	// print pi
	fmt.Println(math.Pi)

	// add two numbers
	fmt.Println(add(24, 12))

	// swap two words
	a, aa := swap("Hello", "Phoenix")
	fmt.Println(a, aa)

	// returns the split functionality
	fmt.Println(split(17))

	// var i int
	// // 0 false false false
	// fmt.Println(i, c, python, java)

	// // 1 2 true false NO!
	// var c, python, java = true,
	// 	false, "NO!"
	// fmt.Println(n, j, c, python, java)

	// name := "Phoenix"
	// age := 25
	// active := true
	// This is equivalent to:
	// var name string = "Phoenix"
	// var age int = 25
	// var active bool = true

	var p, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"
	// 1 2 3 true false no!
	fmt.Println(p, j, k, c, python, java)

	// Complex numbers
	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)
	// Type: bool Value: false
	// Type: uint64 Value: 18446744073709551615
	// Type: complex128 Value: (2+3i)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// Variables declared without an explicit initial value are given their zero value.
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
