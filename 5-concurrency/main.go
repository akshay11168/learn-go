package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/tour/tree"
)

func main() {

	fmt.Println("Learning about go routines")
	// goRoutines()

	// channels()
	// bufferedChannels()
	// rangeAndClose()

	// selectStatment()

	// defaultSelection()

	excercise()

	// syncMutex()
}

func goRoutines() {

	// Go Routines are lightweight thread managed by Go run time.

	// go f(x,y,z) starts a new go routine.
	// The evaluation of x,y,z happens in the current Goroutine and the execution of f happens in the new Goroutine

	// Go routines run in the same address space so access to the shared memory must be synchronized.
	// The Sycn package provides useful primitives although you won't need them much in Go as there are other primitives.

	go say("from go routine")
	say("normal run")

}

func say(uuid string) {

	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("this is loop", uuid, i)
	}
}

func sum(s []int, c chan int) {
	sum := 0

	for _, v := range s {
		sum += v
	}
	c <- sum

}

func channels() {
	// channels are a typed conduit through which we can send or recieve values with the channel operator <-

	// ch <- v // Send v to channel ch
	// v := <- ch // Recieve from channel ch and assign value to v
	// (the data flows in the direction of the arrow)

	// Like maps and slices, channels must be created before use
	// ch := make(chan int)

	// By default sends and recieves block until the other side is ready. This allow goroutine to synchronise without explicit locks or condition variables.

	// The example code sums the number in a slice, distributing the work between two goroutines. Once both the

	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c

	fmt.Println(x + y)

}

func bufferedChannels() {
	// Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel:
	// ch := make(chan int, 100)

	// Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.

	ch := make(chan int, 2)

	ch <- 1
	ch <- 2

	// If we overflow the buffered channel we get the below error
	// fatal error: all goroutines are asleep - deadlock!
	// ch <- 3

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func fibonacci(n int, c chan int) {

	x, y := 0, 1

	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func rangeAndClose() {
	// A sender can close a channel to indicate that no more values will be sent.
	// Receivers can test whether a channel has been closed by assigning a second parameter to the receiver expression:

	// v, ok := <-ch

	// ok is false if there are no more values to receive and the channel is closed

	// The loop for i:= range ch receives the values from the channel repeatedly until it is closed.

	// Note: Only the sender should close the channel, Never the reciever. Sending on a closed channel will cause a panic
	// Note: Channels aren't like files; you don't usually need to close them.
	// Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop

	c := make(chan int, 10)
	go fibonacci(50, c)
	for i := range c {
		fmt.Println(i)
	}
}

func fibonacciSelect(c, quit chan int) {

	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("select statment quits!")
			return
		}
	}
}

func selectStatment() {

	// The select Statement lets a goroutine wait on multiple communication operations
	// A select blocks until one of its cases can run, then it executes that case. It chooses one at the random if multiple are ready

	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacciSelect(c, quit)
}

func defaultSelection() {
	// The default case in a select is run if no other case is ready

	// Use a default case to try a send or receive without blocking :

	// select {
	// case i := <-c:
	// 	// use i
	// default:
	// 	// receiving from c would block
	// }

	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-boom:
			fmt.Println("boom")
			return
		default:
			fmt.Println("this is default")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func excercise() {

	// fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println("making crawler")
}

func Same(t1, t2 *tree.Tree) bool {

	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if v1 != v2 || ok1 != ok2 {
			return false
		}

		if !ok1 {
			break
		}
	}

	return true
}

func Walk(t *tree.Tree, ch chan int) {
	defer close(ch) // <- closes the channel when this function returns
	var walk func(t *tree.Tree)
	walk = func(t *tree.Tree) {
		if t == nil {
			return
		}
		walk(t.Left)
		ch <- t.Value
		walk(t.Right)
	}
	walk(t)
}

// Is safe to use concurrently
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Increaments the counter for the give key.
func (c *SafeCounter) Inc(key string) {

	// Locks so that only one go routine can access the map c.v at one time
	c.mu.Lock()
	defer c.mu.Unlock()
	c.v[key]++
}

func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

func syncMutex() {

	// But what if we don't need communication?
	// What if we want to make sure only one goroutine can access a variable at a time to avoid conflicts

	// This concept is called mutual exclusion, and the conventional name for the data structure that provides it is mutex

	// Go standard library provides mutual exclusion with sync.Mutex and its two methods:
	// 1. Lock
	// 2. Unlock

	// We can define a block of code to be executed in mutual exclusion by surrounding it with a call to Lock and Unlock as shown on the INC method

	// We can also use defer to ensure the mutex will be unlocked as in the value method

	c := SafeCounter{v: make(map[string]int)}

	for i := 0; i < 222; i++ {
		go c.Inc("a")
	}

	time.Sleep(time.Second)

	fmt.Println(c.Value("a"))
}
