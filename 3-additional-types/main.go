package main

import (
	"fmt"
	"math"
	"strings"

	"golang.org/x/tour/pic"
)

func main() {

	// pointers()
	// structs()
	// pointersToStructs()

	// structLiterals()

	// arrays()

	// slices()

	// sliceLiteral()
	// sliceDefault()
	// sliceNil()

	// sliceWithMake()
	// slicesOfSlices()
	// appendingToSlice()

	// rangeFunction()

	// maps()

	// mapsMutate()
	challenge()

	// functionValues()

	// functionClosure()

}

func pointers() {

	// pointer holds memory to the address of a value

	// The type *T is a pointer to a T value. Its zero value is nil
	var p *int
	fmt.Println(p)

	// The & operation generates a pointer to a operand
	i := 10
	p = &i

	// * operater denotes the pointer's underlying value
	fmt.Println("value of p is", p)
	fmt.Println("value of *p is", *p)
	fmt.Println("this is know as 'dereferencing' or 'indirecting'")

	// this is know as "dereferencing" or "indirecting"
}

type vertex struct {
	X int
	Y int
}

func structs() {
	// A struct is a collection of fields

	fmt.Println(vertex{1, 2})

	// struct fields are accessed using a dot notation

	v := vertex{1, 2}

	fmt.Println("struct fields are accessed using a dot notation,like v.Y ", v.X)
}

func pointersToStructs() {
	// Struct fields can be accessed through a struct pointer.

	// To access the field X of a struct when we have the struct pointer p we could write (*p).X. However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.

	v := vertex{10, 20}

	p := &v

	fmt.Println("this is a pointer p to struct vertex ", p)
	fmt.Println("the value of p.Y is", p.X)
	fmt.Println("for accessing the fields inside the struct pointer there is no need for explicit dereferencing ")
}

func structLiterals() {

	type vertex struct {
		X, Y int
	}

	var (
		v1 = vertex{1, 2}  // has type vertex
		v2 = vertex{X: 1}  // Y :0 is implicit
		v3 = vertex{}      //X:0 and Y:0
		p  = &vertex{3, 4} //has type *vertex
	)

	fmt.Println(v1, v2, v3, p)
	fmt.Printf("v1 has value %v and type %T for v1 = vertex{1, 2} \n", v1, v1)
	fmt.Printf("v2 has value %v and type %T for v2 = vertex{X: 1} \n", v2, v2)
	fmt.Printf("v3 has value %v and type %T for v3 = vertex{} \n", v3, v3)
	fmt.Printf("p has value %v and type %T for p = &vertex{3, 4} \n", p, p)
}

func arrays() {
	// The type [n]T is an array of n values of type T

	var a [10]int

	// An arrays length is part of its type so arrays cannot be resized. This seems limiting, but don't worry; Go provides a convenient way of working with arrays
	fmt.Println(a)

	var s [10]string
	s[1] = "hello"

	fmt.Println("value of s is", s, "& value of s[1] is ", s[1])

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func slices() {
	// An array has a fixed size, where are slice has a dynamic sizing.

	// the type []T is a slice with elements of type T

	var x []int

	// A slice is formed by specifiying the two indices lower bound and upper bound, seperated by colon;
	// a[low : high]
	// This selects a half-open range which includes the first element, but excludes the last one.
	// The following expression creates a slice which includes elements 1 through 3 of a:
	// a[1:4]

	fmt.Println("s", x)

	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)

	// Slice are like references to array
	// A slice doesn't store any data. It describes a section of an underlying array
	// changing the element of a slice modifies the element of the underlying array
	// If other slice sharing is the same array then the value is changed for both of them

	var fullName [14]string = [14]string{"a", "k", "s", "h", "a", "y", " ", "b", "i", "r", "a", "d", "a", "r"}
	fmt.Printf("variable fullName has value %v and is of type %T \n", fullName, fullName)
	var firstName []string = fullName[0:6]
	fmt.Printf("variable firstName has value %v and is of type %T \n", firstName, firstName)
	var lastName []string = fullName[7:]
	fmt.Printf("variable lastName has value %v and is of type %T \n", lastName, lastName)

	var random [4]string = [4]string{"A", "B", "C", "D"}
	fmt.Println(random)
	aSlice := random[:3]
	bSlice := random[1:]

	fmt.Printf("%v , %v \n", aSlice, bSlice)

	aSlice[2] = "X"

	fmt.Printf("%v , %v \n", aSlice, bSlice)

}

func sliceLiteral() {

	// A slice literal is like an array literal without the length

	// Below is array literal
	// [3]bool{true,false,true}
	// And slice creates the same array as above, then builds a slice that reference it
	// []bool{true,false,true}

	var q []int = []int{2, 3, 5, 7, 11, 13}
	fmt.Printf("value of slice q is %v and type is %T \n", q, q)

	r := []bool{true, false, true, false, true}
	fmt.Printf("value of slice r is %v and type is %T \n", r, r)

	s := []struct {
		t bool
		v int
	}{
		{true, 2},
		{false, 0},
	}

	fmt.Printf("value of slice s is %v and type is %T \n", s, s)
}

func sliceDefault() {
	// When slicing an array we can omit the lower bound or higher bound to use their default
	// low bound = 0
	// high bound = length of slice

	// For the array
	// 	For the array
	// var a [10]int
	// these slice expressions are equivalent:
	// a[0:10]
	// a[:10]
	// a[0:]
	// a[:]

	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s, len(s), cap(s)) // [3 5 7] 3 5

	s = s[:2]
	fmt.Println(s, len(s), cap(s)) // [3 5] 2 5

	s = s[1:]
	fmt.Println(s, len(s), cap(s)) // [5] 1 4

	// Slice Length - The length of slice is the number of elements it contains
	// Slice Capacity - The capacity of the slice is the number of elements in the underlying array, counting from the first element in the slice.

}
func sliceNil() {
	// Zero value of a slice is nil
	// A nil slice has a length and capacity of 0 and has no underlying array.

	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("this is a nil slice")
	}
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func sliceWithMake() {

	// slice can be created with builtin make function. this is how we create a dynamically sized arrays
	// The make function allocates a zeroed array and returns a slice that refers to that array

	a := make([]int, 5)
	printSlice("a", a)

	// To specify the capacity pass a third argument to make
	b := make([]int, 5, 10)

	printSlice("b", b)
}

func slicesOfSlices() {
	// A slices can contain any type including other slices
	s := [][]int{[]int{1}, []int{2}}
	// printSlice("S", s)
	fmt.Println(s)

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func appendingToSlice() {
	// TODO Go provides a built in append function : https://pkg.go.dev/builtin#append
	// TODO https://go.dev/blog/slices-intro
	// func append(s []T, vs ...T) []T
	// First parameter of append function is a slice of type T
	// Second parameter of append function is values of T
	// Resulting value of append is a slice containing all the elements of original slice plus the value provided.
	// If the backing array of s is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array.

	var s []int
	printSliceString(s)

	// append works on nil slices.
	s = append(s, 0)
	printSliceString(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSliceString(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSliceString(s)

}

func printSliceString(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func rangeFunction() {
	// The Range form of the for loop iterates over a slice or map
	// When ranging over a slice two values are returned for each iteration.
	// First - Index
	// Second - Copy of the element at that index.

	var pow = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i, v := range pow {
		fmt.Printf("at index %v value is %v \n", i, v)
	}

	// We can skip the index of value by assigning _
	for _, v := range pow {
		fmt.Printf("each value is %v \n", v)
	}

	for i, _ := range pow {
		fmt.Printf("each index is %v \n", i)
	}
	// If we want only index we can omit the second variable
	for i := range pow {
		fmt.Printf("each index without value is %v \n", i)
	}
}

func challenge() {

	// one()
	// two()
	three()
}

func one() {

	pic.Show(createPic)

}

func createPic(dx, dy int) [][]uint8 {

	fmt.Println(dx)
	fmt.Println(dy)

	slice := make([][]uint8, dy)

	for i := range slice {
		slice[i] = make([]uint8, dx)
	}

	return slice
	// make([])
}

func maps() {
	// A Map Maps keys to values
	// Zero value of map is nil, A nil map has no keys, nor can keys be added.
	// The make function returns a map of give type, initialized and ready to use

	type vertex struct {
		x int
		y int
	}

	var m map[string]vertex = map[string]vertex{
		"def": vertex{1, 2},
		"ghi": vertex{3, 4},
		// If the top level type is just a type name, you can omit it from the elements of the literal
		"jkl": {2, 3},
	}
	fmt.Println(m)

	m = make(map[string]vertex)
	fmt.Println(m)

	m["abc"] = vertex{1, 3}
	fmt.Println(m)
}

func mapsMutate() {
	var m map[string]int = make(map[string]int)

	// insert a new element in map
	m["a"] = 0
	m["A"] = 0
	m["b"] = 2
	// update a element in map
	m["a"] = 1

	fmt.Println(m)

	// retrieving an element from map
	first := m["a"]
	fmt.Println(first)

	// Deleting and element
	delete(m, "A")
	fmt.Println(m)

	// Test if the element is present in a map
	// Key is in m if ok is true, If not ok is false
	// If key is not present in the map then the elem value will be zero value of the type of element of the map
	elem, ok := m["A"]
	fmt.Println(elem, ok)

}

func two() map[string]int {

	var wordCount map[string]int = make(map[string]int)

	var s string = "asfa af as fas fas fas f asf sasaf asfa sf asfa"

	for _, v := range strings.Fields(s) {
		// fmt.Println(i, v)
		wordCount[v] = wordCount[v] + 1
	}
	// fmt.Println(wordCount)
	return wordCount
}

func functionValues() {
	// Functions are values too. they can be passed around like other values
	// Function values may be used as function argument and return values

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(3, 4))
	fmt.Println(compute(hypot))
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(10, 20)
}

// adder function returns a closure
func adder() func(int) int {
	sum := 0

	return func(x int) int {
		sum = sum + x
		return sum
	}
}

func functionClosure() {

	// A Go func can be closure. A closure is a function value that references value from outside its body.
	// The function may access and assign to the referenced variables; In this sense  the function is "bound" to the varaibles

	// adder function returns a closure. Each closure is bound to its own sum variable.
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func fibonacci() func() int {

	current := 0
	previous := 0
	return func() int {
		value := current

		if current == 0 && previous == 0 {
			current = current + 1
		} else {
			temp := previous
			previous = current
			current = previous + temp
		}
		// else {
		// 	sum = sum + sum
		// }

		return value
	}
}
func three() {
	f := fibonacci()
	for i := 0; i < 50; i++ {
		fmt.Println(f())
	}
}
