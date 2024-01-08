// Every GO Program is made of package
// Program always starts running in package main
package main

// The code groups the import into a paraenthesis, Also know as "factored" import statement
import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

// var statement declares a list of variable; as in function arguments list, the type is last
var global, universal, extensive bool

// var declaration can include initializer, one per variable
var a, b, c int = 1, 2, 3

func main() {
	// var can be function level or Package level
	var local, confined, restricted bool
	var i, j, k = 1, 2, "true"

	fmt.Println("Hello World")

	// Uses time package to get current time.
	fmt.Println("Current Time is", time.Now())

	// Package name is by default the last element of the path.
	fmt.Println("Random Number generated is", rand.Intn(10))

	// In GO a name is exported if is begins with Capital Letter, math.pi will not be exported but math.Pi will be
	fmt.Println("Value of Math constant PI is", math.Pi)

	// TO Understand the Go type, read the below blog to understand why the syntax style choosen was this
	// https://go.dev/blog/declaration-syntax
	// https://c-faq.com/decl/spiral.anderson.html

	// Arguments are the values send to a function when it is called.
	fmt.Println("1 + 2 =", addition(1, 2))

	fmt.Println("2 - 1 =", substract(2, 1))

	// Getting two values from the function Swap
	variable_one, variable_two := swap("one", "two")
	fmt.Println("This function returns two value, swap('one','two')", variable_one, ",", variable_two)

	// Spliting the number using Named return values
	split_one, split_two := split(10)
	fmt.Println("Splitting the number(10) using Named Return Value (uses naked return statments)", split_one, split_two)

	fmt.Println("A var variable can be function level or package level, local variables are : ", local, confined, restricted, "global variables are", global, extensive, universal)

	fmt.Println("These varibales are declared and initialised globally with explicit type declaration", a, b, c)

	fmt.Println("These variables are declared and initialised locally without explicit type declaration", i, j, k)

	// Short Variable declaration
	// Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type
	l, m := 10, 20
	fmt.Println("The := short assignment statement can be used in place of a var declaration with implicit type inside a function ; Sample variable values are ", l, m)
	// Outside of function, every keyword begins with a keyword (var, func and so on) and so the := construct is not available

	types()
}

// Parameter are the variables listed inside parantheses of a function definition
func addition(a int, b int) int {
	return a + b
}

// When two or more consecutive named function parameters share a type. we can omit the type from all but the last
func substract(a, b int) int {
	return a - b
}

// A function can return a number of results
func swap(x, y string) (string, string) {
	return y, x
}

// Go's return values may be named if so they are treated as variable defined at the top
// These names should be used to documentthe meaning of the return values
func split(total int) (a, b int) {

	a = total * 4 / 9
	b = total - a

	// a return statement without arguments returns the named return values. This is also know as "Naked" return
	return

	// Naked returns should be used only for short functions as they harm the readability in the longer function
}

func types() {

	// Go basic types :
	// bool
	// string
	// int  int8  int16  int32  int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte // alias for uint8
	// rune // alias for int32
	// 	// represents a Unicode code point
	// float32 float64
	// complex64 complex128

	// variable declaration may be "factored" into blocks as with import statements
	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	zero_values_types()
}

func zero_values_types() {

	// variables declared without an explicit initial value are given their zero value

	// 0 for numeric types,
	// false for the boolean type, and
	// "" (the empty string) for strings.

	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	typeConversion()
}

// TODO: Continued 8th Jan 2024
func typeConversion() {

}
