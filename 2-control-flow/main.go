package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {
	loops()

	conditions()

	switchConditions()

	deferStatements()
	challenges()

}

func loops() {

	// Go has only on looping construct "FOR"

	// A basic for loop has three construct
	// 1. init Statement : executed before first execution
	// 2. the condition statement : evaluated before every execution
	// 3. the post statement : executed at the end o every iteration

	// Varaibles declared in the init statement have a scope of the loop
	for i := 1; i < 3; i++ {
		fmt.Println("Current loop is ", i)
	}

	// for Fo's while
	count := 0

	for count < 2 {
		fmt.Println("loop without init and post statment with count value is", count)
		count++
	}

	// // FOREVER
	// // if we omit the loop condition it becomes infinite loop
	// for {
	// fmt.Println("I will run forever")
	// }
	fmt.Println("I will run forever if above line is commented")
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func challenges() {

	fmt.Println("trying out suqare root challenge")

	fmt.Println(squareRootChallenge(225))
}

func conditions() {
	// Go's If statement's expression need not be covered with () but {} are required

	x := -1
	if x < 0 {
		fmt.Println("X is less than zero based on avobe condition")
	}

	// square roots
	fmt.Printf("square root of 2 is %v and square root of -2 is %v \n", sqrt(2), sqrt(-2))
	fmt.Println("power of 2 to 2 is", pow(2, 2, 10))
	fmt.Println(pow(5, 5, 10))
}

func pow(x, n, max float64) float64 {

	// the if condition can start with a short statement to execute before the condition
	// the variables declared inside this condition has the scope until the end of if
	if v := math.Pow(x, n); v < max {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, max)
	}

	return max
}

func squareRootChallenge(num float64) float64 {

	var z float64 = 1
	var x float64
	count := 0
	for ; count < 20; x = z {
		z -= (z*z - num) / (2 * z)
		if x == z {
			return z
		}
		fmt.Printf("on loop %v value of z is %v \n", count, z)
		count++
	}
	return z
}

func switchConditions() {

	// switch case is a shorter way of writting a sequence of if-else condition
	// It has not break conditions unlike other languages. break conditions are by default there on each case
	// switch cases evaluate cases from top to bottom
	guessMyOs()

	guessNextSaturday()

	trueSwitchCase()
}

func guessMyOs() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func guessNextSaturday() {

	today := time.Now().Day()

	fmt.Printf("Next Saturday when? ")
	switch today {
	case today + 0:
		fmt.Println("today is saturday")
	case today + 1:
		fmt.Println("tommorow is saturday")
	case today + 2:
		fmt.Println("Day after tommorow is saturday")
	default:
		fmt.Println("too far Don't think much !!!")
	}
}

func trueSwitchCase() {
	fmt.Println("This switch case has no condition, It is a good construct to write if then else conditions")

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func deferStatements() {

	deferOne()

	fmt.Println("now completed the deferred section ! ", deferStack())
}

func deferOne() {
	//  A defer statment defers the execution of a function until the surrounding function returns

	// The functions deferred call arguments are evaluated immediately but the function call is not executed until the surrounding function returns
	defer fmt.Println("deferOne : first ")
	fmt.Println("deferOne : second")
}

func deferStack() string {
	fmt.Println("stack of defered functions")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	// Deferred function calls are pushed onto a stack. When a function returns its deferred calls are executed in last-in first-out order
	// More at TODO : https://go.dev/blog/defer-panic-and-recover
	return fmt.Sprint("done! now read this https://go.dev/blog/defer-panic-and-recover")
}
