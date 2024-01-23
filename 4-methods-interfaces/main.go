package main

import (
	"fmt"
	"image"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type I interface {
	M()
}

type T struct {
	S string
}

func main() {

	fmt.Println("learning about methods and interfaces")

	methods()
	pointerReciever()

	methodAndPointerIndirection()
	interfaces()
	interfacesImplicit()

	interfaceValuesWithNil()

	nilInterface()

	emptyInterface()

	typeAssertions()

	typeSwitches()

	stringersIntro()

	stringersExcercise()

	Errors()
	ErrorsExcercise()

	Readers()
	rot13ReaderExcercise()

	imageFunctions()
}

type vertex struct {
	x float64
	y float64
}

func (v vertex) abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func absFunc(v vertex) float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func methods() {
	// Go doesn't have classes. However we can define methods on types.

	// A method is a function with a special reciever argument

	// A reciever appears in its own argument list between the func keyword and the method name

	v := vertex{1, 2}
	fmt.Println(v.abs())

	// Remember : a method is just a function with a reciever argument
	// Here absFunc is written as a regular function  with no change in functionality

	fmt.Println(absFunc(v))

	a := MyFloat(1)
	a.Abs()
}

// We can declar a method on non-struct types, too.

// We can only declare a method with a reciever whose type is defined in the same package as the method.
// You cannot declare a method with a reciever whose type is defined in another package (which includes the built-in types such as int).

type MyFloat float64

func (f MyFloat) Abs() {
	fmt.Println(f)
}

func pointerReciever() {
	// We can declare methods with pointer reciever
	// This means the receiver type has the literal syntax *T for some type T. (Also T cannot itself be a pointer such as *int ).

	// For example, the scale method below is defined on *vertex

	v1 := vertex{1, 2}
	fmt.Println("value before scaling", v1)
	v1.scale(5)
	fmt.Println("value after scaling", v1)

	v2 := vertex{2, 5}
	fmt.Println("value before scaling", v2)
	scaleFunction(&v2, 5)
	fmt.Println("value after scaling", v2)
}

// This * pointer notation makes sure that we can change the same variable we are propagated
func (v *vertex) scale(f float64) {

	v.x = v.x * f
	v.y = v.y * f

}
func scaleFunction(v *vertex, f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func methodAndPointerIndirection() {

	// Comparing the two programs, We might notice that the function with a pointer argument must take a pointer

	v := vertex{1, 2}
	// scaleFunction(v,5) // Error as v must be a pointer
	scaleFunction(&v, 5) // ok!

	// While Method with pointer reciever can take either a value or a pointer as the reciever when they are called
	v.scale(5)
	p := &v
	p.scale(2)

	// For the statement v.scale(), even though v is a value not a pointer,
	//  the method with the pointer reciever is called  automatically.
	// That is, as a convience, Go interprets the statment, v.scale(5) as (&v).scale(5) since the Scale method has a pointer reciever

	// The equivalent things also happens in reverse direction

	// Function that take a value argument must take a value  of that specific type.
	var b vertex
	fmt.Println(absFunc(b)) // OK
	// fmt.Println(absFunc(&b)) // Compile error!

	// while methods with the value reciever take either a value or a pointer as the reciever when they are called.
	var a vertex = vertex{2, 2}
	fmt.Println(a.abs())
	pa := &a
	fmt.Println(pa.abs())

	// Choosing a value or pointer receiver
	// There are two reasons for using pointer receiver
	// 1. It gives function capability for change the value of the variable itself. that is value is changed at the reciever itself.
	// 2. To avoid copying the value on each method call. This can be more efficient if reciever is a larger struct.

	// In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both. (We'll see why over the next few pages.)

}

type abser interface {
	absolute() float64
}

func (f MyFloat) absolute() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *vertex) absolute() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func interfaces() {
	// Interface type is defined as a set of method signatures.
	// A value of interface type can hold any value that implements those methods.

	// Note: There is an error in the example code on line 22. Vertex (the value type) doesn't implement Abser because the Abs method is defined only on *Vertex (the pointer type)

	var a abser
	f := MyFloat(2)
	v := vertex{3, 4}

	fmt.Println(a)
	fmt.Println(f)
	fmt.Println(v)
	// a = v
	a = f
	a = &v

	// fmt.Println(a.absolute())
}

type Im interface {
	m()
}

func (t T) m() {
	return
}

func interfacesImplicit() {

	// A type implements an interface by implementing its methods. There is no explicit declaration of intent, no implements keyword.
	// Implicit interfaces decouple the definition of an interface from its implementation, Which would then appear in any package without prearrangement

	var a Im = T{"saafa"}

	fmt.Println("a", a)

	// Under the hood, interface values can be thought of as a tuple of a value and a concrete type: (value, type)

	// An interface value holds a value of a specific underlying concrete type.
	// Calling a method on an interface value executes the method of the same name on its underlying type.

	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func interfaceValuesWithNil() {

	// If the concrete value inside the interface is nil, the method will be called a nil receiver
	// In Go it is common to write methods that gracefully handle being called with a nil receiver.

	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
	// Note that an interface value that holds a nil concrete value is itself non-nil

}

func nilInterface() {

	// A nil interface values holds neither value nor concrete type.

	// Calling a method with nil interface is a run time error because there is no type inside the interface tuple to indicate which concrete method to call.

	var i I
	describe(i)

	// error := i.M()

	// fmt.Println("this is the erroro", error)
}

func emptyInterface() {

	// The interface types that specifies zero method is know as the empty interface.
	// interface{}

	// An empty interface may hold values of any type. (every type implements atlease zero method)

	// Empty interface are used by code that handles values of unknow type. For example, fmt.Print takes any number of arguments of type interface{}

	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "abc"
	describe(i)
}

func typeAssertions() {
	// Type assertion provides access to an interface value's underlying concrete values
	// t := i.(T)

	// This statement asserts that the interface value i holds the concrete type T and assigns the  underlying value T to the variable t
	// If i doesn't hold a value T, The statment will trigger panic

	// To test whether an interface value holds a specific type, a type assertion can return two values:
	// the underlying value and a boolean value that return whether the assertion succeeded.
	// t, ok := i.(T)

	// If i hold a value T, then t will be the underlying value and ok will be true.
	// if not ok will be false. and t will be Zero value of type T, and no panic occurs.

	// Note the similarity between this syntax and that of reading from a map.

	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	a, ok := i.(string)
	fmt.Println(a, ok)

	b, ok := i.(int)
	fmt.Println(b, ok)

	// this will panic
	// c := i.(int)
	// fmt.Println(c)

}

func do(i interface{}) {

	//  The declaration in a type switch has the same syntax as the type assertion i.(T), but the specific type T is replaced with the keyword type.

	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}

	// switch v := i.(type) {
	// case T:
	// 	// here v has type T
	// case S:
	// 	// here v has type S
	// default:
	// 	// no match; here v has the same type as i
	// }

	// This switch statement tests whether the interface value i holds a value of type T or S.
	// In each of the T and S cases, the variable v will be of type T or S respectively and hold the value held by i.
	// In the default case (where there is no match), the variable v is of the same interface type and value as i.
}
func typeSwitches() {
	// A type switch is construct that permits se

	// A type switch is like a regular switch statement, but the cases in type switch specify types (not values)
	// and those values are compared against the type of the value held by the give interface value

	do(21)
	do("hello")
	do(true)
}

type Stringer interface {
	String() string
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func stringersIntro() {
	// One of the most ubiquitous interface is Stringer defined by the fmt package.

	// type Stringer interface {
	// 	String() string
	// }

	// A Stringer is a type that can describe itself as a string.
	// The fmt package (and many others) look for this interface to print values.

	p := Person{"akshay", 26}

	fmt.Println(p)

}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func stringersExcercise() {

	hosts := map[string]IPAddr{
		"localhost": {127, 0, 0, 1},
		"test":      {127, 2, 0, 1},
	}

	for name, ip := range hosts {
		fmt.Printf("%v : %v \n", name, ip)
	}

}

type MyError struct {
	When  time.Time
	Where string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s ", e.When, e.Where)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func Errors() {
	// A Go program express error state with error values

	// The error type is a built in interface similar to fmt.Stringer

	// type error interface {
	// 	Error() string
	// }

	// As with fmt.stringer the fmt package looks for the error interface when printing values.

	// Function often return an error value and calling code should handle errors by testing whether the error equals nil

	i, err := strconv.Atoi("42")
	if err != nil {
		fmt.Printf("couldn't convert number: %v\n", err)
		return
	}
	fmt.Println("Converted integer:", i)

	// A nil error denotes success; a non-nil error denotes failure.

	if err := run(); err != nil {
		fmt.Println(err)
	}
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// TODO:
	// Note: A call from fmt.Sprint(e) inside the Error method will send the program into an infinite loop.
	// We can avoid this by converting e first: fmt.Sprint(float64(e)). Why?

	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < -1 {

		// TODO: WHY
		// Change your Sqrt function to return an ErrNegativeSqrt value when given a negative number.
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}
func ErrorsExcercise() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

func Readers() {
	// The io package specifies the io.Reader interface, which represents the read end of a stream of data

	// The GO standard library contains many implementations of this interface, including files, network connections, compressors, ciphers and others
	// TODO https://cs.opensource.google/search?q=Read%5C(%5Cw%2B%5Cs%5C%5B%5C%5Dbyte%5C)&ss=go%2Fgo

	// The io.Reader interface has a Read method
	// func (T) Read(b []byte) (n int, err error)

	// Read populates the given byte slice with data and returns the number of bytes populated and an error value.
	// It returns io.EOF error when the stream ends

	// This code creates a string.Reader and consumes its output 8 bytes at a time.

	r := strings.NewReader("Hello World")

	b := make([]byte, 8)

	for {
		n, err := r.Read(b)
		if err == io.EOF {
			fmt.Println(err)
			break
		}
		fmt.Println("here", n)
	}
}

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {

	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

type rot13Reader struct {
	r io.Reader
}

func rot13(x byte) byte {

	switch {
	case x >= 65 && x <= 77:
		fallthrough
	case x >= 97 && x <= 109:
		x = x + 13
	case x >= 78 && x <= 90:
		fallthrough
	case x >= 110 && x <= 122:
		x = x - 13
	}
	return x

}
func (r rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)

	for i := range b {
		b[i] = rot13(b[i])
	}

	return n, err
}

func rot13ReaderExcercise() {

	// A common pattern is an io.Reader that wraps another io.Reader, modifying the stream in some way.
	// For example, the gzip.NewReader function takes an io.Reader (a stream of compressed data) and returns a *gzip.Reader that also implements io.Reader (a stream of the decompressed data).

	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)

	fmt.Println()
}
func imageFunctions() {

	// Package image defined Image interface

	// package image
	//	type Image interface {
	//	    ColorModel() color.Model
	//	    Bounds() Rectangle
	//	    At(x, y int) color.Color
	//	}

	// Note : the Rectangle return value of the bounds method is actually an image.Rectangle, as the declaration is inside package image.

	// The color.Color and color.Model types are also interfaces, but we'll ignore that by using the predefined implementations color.RGBA and color.RGBAModel. These interfaces and types are specified by the image/color package

	m := image.NewRGBA(image.Rect(0, 0, 100, 100))

	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
