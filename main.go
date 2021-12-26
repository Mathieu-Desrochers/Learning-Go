package main

import (
	"bytes"
	"fmt"
	"sort"
	"time"
	"unicode/utf8"
)

func main() {

	// variable declarations
	var number int = 1
	var one, two = 1, 2
	three := 3

	// unused variables produce compilation errors
	fmt.Println(number + one + two + three)

	// arrays have a fixed length
	var array [2]int
	fmt.Printf("array of %v elements\n", len(array))

	// array literals
	_ = [3]int{1, 2, 3}
	_ = [...]int{1, 2, 3, 4}
	_ = [...]int{2: 10, 4: 20}

	// slices have an auto growing length
	// they keep track of an array and its capacity
	var slice []int
	fmt.Printf("slice of %v elements and a capacity for %v\n", len(slice), cap(slice))

	// slice literals
	_ = []int{}
	_ = []int{1, 2, 3, 4}
	_ = []int{2: 10, 4: 20}

	// appending values
	// can reallocate the array to a bigger location
	// must be recaptured
	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3)
	fmt.Printf("appended slice %v\n", slice)

	// selecting values
	fmt.Printf("selected slice %v\n", array[1:])

	// modifying values
	slice[0] = 10
	slice[1] = 20
	slice[2] = 30
	fmt.Printf("modified slice %v\n", slice)

	// removing values
	copy(slice[1:], slice[2:])
	slice = slice[:len(slice)-1]
	fmt.Printf("removed slice %v\n", slice)

	// slices can be built with a
	// predefined length and capacity
	slice = make([]int, 5, 1000)
	fmt.Printf("slice of %v elements and a capacity for %v\n", len(slice), cap(slice))

	// selected slices are not stable
	// changes in the source can be seen
	// unless the source gets reallocated
	// probably make a copy
	source := []int{1}
	selectedSlice := source[:]
	source[0] = 2
	fmt.Printf("selected slice %v\n", selectedSlice)
	source = append(source, 0)
	source[0] = 3
	fmt.Printf("selected slice %v\n", selectedSlice)

	// maps are hash tables
	var nameById = make(map[int]string)

	// map literals
	_ = map[int]string{}
	_ = map[int]string{1: "Alice", 2: "Bob"}

	// setting values
	nameById[100] = "Alice"
	nameById[200] = "Bob"
	nameById[300] = "Carl"

	// looking up values
	if name, ok := nameById[100]; ok {
		fmt.Printf("name: %v\n", name)
	}

	// iterating over values
	// order is not guaranteed
	for id, name := range nameById {
		fmt.Printf("id: %v, name: %v\n", id, name)
	}

	// removing values
	delete(nameById, 300)

	// strings are immutable sequences of bytes
	greek := "some greek: Τη γλώσσα μου έδωσαν"

	// the index operation returns a byte
	fmt.Println(greek[0])

	// the substring operation returns a string
	fmt.Println(greek[5:10])

	// strings can be decoded as bytes
	greekBytes := []byte(greek)
	fmt.Printf("greek decoded as bytes: %v\n", greekBytes)

	// or as utf8 unicode code points
	// these are named runes and are int32
	greekRunes := []rune(greek)
	fmt.Printf("greek decoded as runes: %v\n", greekRunes)

	// fancy decoding is required to
	// index the runes inside a string
	rune, bytesCount := utf8.DecodeRuneInString(greek[14:])
	runesCount := utf8.RuneCountInString(greek)
	fmt.Printf("found rune %c spanning %v bytes\n", rune, bytesCount)
	fmt.Printf("found %v runes\n", runesCount)

	// iterating is done over runes
	for range greek {
	}

	// efficient string building using a buffer
	var buffer bytes.Buffer
	buffer.WriteByte('a')
	buffer.WriteRune('λ')
	buffer.WriteString("yeah")
	fmt.Println(buffer.String())

	// structure definitions
	type Employee struct {
		EmployeeID int
		FirstName  string
		LastName   string
	}

	// structure literals
	_ = Employee{1, "Alice", "Alisson"}
	_ = Employee{FirstName: "Alice"}

	// structure allocations
	_ = new(Employee)
	_ = &Employee{2, "Bob", "Bobson"}
	_ = &Employee{FirstName: "Bob"}

	// accessing fields
	var employee Employee = Employee{FirstName: "A"}
	fmt.Printf("employee first name: %v\n", employee.FirstName)

	// same notation with pointers
	var employeePointer *Employee = &employee
	fmt.Printf("employee first name: %v\n", employeePointer.FirstName)

	// structures are passed by value
	// but are primarily used with pointers
	type Team struct {
		Manager   *Employee
		Employees []*Employee
	}

	// anonymous structures
	var point struct{ X, Y int }
	point.X = 100

	// anonymous structure literals
	_ = struct{ X, Y, Z int }{X: 1, Y: 2, Z: 3}

	// named types
	type ShoeSize int
	var _ ShoeSize = ShoeSize(14)

	// something like an enum
	type Flavor int32
	const (
		Vanilla Flavor = iota
		Chocolate
		Pistachios
	)
	var bestFlavor Flavor = Chocolate
	fmt.Printf("bestFlavor: %v\n", bestFlavor)

	// see you later
	later()
}

// function signatures
func noReturn() {
}
func oneReturn() bool {
	return false
}
func multipleReturns() (bool, int) {
	return true, 25
}
func bareReturns() (x, y int) {
	x = 1
	y = 2
	return
}

// there is no tail call optimization
// but we get auto-growing stacks
func recurse(x int) {
	if x < 1000 {
		recurse(x + 1)
	}
}

// error handling
func ooops() error {
	return fmt.Errorf("damn thing exploded")
}
func errorPropagation() error {
	err := ooops()
	if err != nil {
		return err
	}
	return nil
}
func errorWithContext(color string) error {
	err := ooops()
	if err != nil {
		return fmt.Errorf("while trying to paint %s: %v", color, err)
	}
	return nil
}

// functions as values
func addNumbers(x, y int) int {
	return x + y
}

func later() {

	// returns
	noReturn()
	_ = oneReturn()
	_, _ = multipleReturns()
	_, _ = bareReturns()

	// functions as values
	var functionAsValue func(int, int) int = addNumbers
	fmt.Println(functionAsValue(1, 2))

	// anonymous functions
	plusOne := func(x int) int { return x + 1 }
	fmt.Println(plusOne(1))

	// closures
	someNumber := 25
	plusTwo := func() int { return someNumber + 2 }
	fmt.Println(plusTwo())

	// but by reference
	someNumber = 50
	fmt.Println(plusTwo())

	// leading to weird patterns
	// where closed values need to be copied
	plusThreeNumber := someNumber
	plusThree := func() int { return plusThreeNumber + 3 }
	someNumber = 75
	fmt.Println(plusThree())

	// variadic functions
	bigCompute := func(values ...int) int {
		return len(values)
	}
	bigComputeValues := []int{1, 2, 3}
	fmt.Println(bigCompute(1, 2, 3))
	fmt.Println(bigCompute(bigComputeValues...))

	// deferred function calls
	// called when the function exits
	doStuff := func() {
		fmt.Println("enter")
		defer fmt.Println("executed")
		{
			defer fmt.Println("not when a block exits")
		}
		fmt.Println("exit")
	}
	doStuff()

	// panicking
	ohNoes := func() {
		panic("we are screwed")
	}

	// recovering
	keepCalm := func() {
		defer func() {
			whatNow := recover()
			fmt.Println(whatNow)
		}()
		ohNoes()
		fmt.Println("too bad won't execute")
	}
	keepCalm()

	laterr()
}

type Animal struct {
	LegsCount int
}

// methods are attached to a receiver type
func (a *Animal) CanQuack() bool {
	return false
}

// receivers are primarily pointers
// to allow state mutations
func (a *Animal) GrowLeg() {
	a.LegsCount++
}

func laterr() {

	// methods
	animal := &Animal{4}
	fmt.Println(animal.CanQuack())

	// structure embedding
	type Dog struct {
		Animal
		GoodBoyName string
	}

	// the structure gains all
	// the members of the embedded one
	fido := &Dog{Animal{4}, "Fido"}
	fmt.Printf("legs count: %v\n", fido.LegsCount)
	fmt.Printf("good boy name: %v\n", fido.GoodBoyName)

	// including its attached methods
	fido.GrowLeg()

	// the embedded structure
	// can be accessed explicitly
	var _ *Animal = &fido.Animal

	// converting from method to a function
	// taking the receiver as first parameter
	methodExpression := (*Animal).GrowLeg
	methodExpression(animal)

	// converting from method to a function
	// with the receiver already bound
	methodValue := animal.GrowLeg
	methodValue()

	laterrr()
}

// encapsulation
// members starting with a lower cased letter
// are only visible inside their package
type Cake struct {
	hugeCaloriesCount int
}

// getters and setters
func (cake *Cake) HugeCaloriesCount() int {
	return cake.hugeCaloriesCount
}
func (cake *Cake) SetHugeCaloriesCount(value int) {
	cake.hugeCaloriesCount = value
}

// interfaces
type Quacker interface {
	Quack(times int)
}

// uses duck typing
// you have the methods you qualify
type Duck struct{}

func (duck *Duck) Quack(times int) {
	for i := 0; i < times; i++ {
		fmt.Printf("quack")
	}
}

func laterrr() {

	// visible inside this package
	var hugeCake = &Cake{100000}
	_ = hugeCake.hugeCaloriesCount

	// any type with a Quack method can be passed
	doTheQuacking := func(quacker Quacker, times int) {
		quacker.Quack(times)
	}
	duck := &Duck{}
	doTheQuacking(duck, 3)

	// the empty interface
	// everyone can play
	var empty interface{}
	empty = false
	empty = 10
	empty = duck
	_ = empty

	// an interface that is nil
	var nilInterface Quacker = nil
	if nilInterface != nil {
		fmt.Println("will not execute")
	}

	// an interface that points to nil
	// never do that
	var nilDuck *Duck = nil
	nilInterface = nilDuck
	if nilInterface != nil {
		fmt.Println("will execute")
	}

	laterrrr()
}

type Cookie struct {
	Size    int
	Flavour string
	Rating  int
}

type CookieSlice []*Cookie

// any type with these
// methods can be sorted
type CookieBySizeSlice []*Cookie

func (x CookieBySizeSlice) Len() int           { return len(x) }
func (x CookieBySizeSlice) Less(i, j int) bool { return x[i].Size < x[j].Size }
func (x CookieBySizeSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// assembling an interface
// from anonymous functions
type FuncSorter struct {
	len  func() int
	less func(i, j int) bool
	swap func(i, j int)
}

func (x *FuncSorter) Len() int           { return x.len() }
func (x *FuncSorter) Less(i, j int) bool { return x.less(i, j) }
func (x *FuncSorter) Swap(i, j int)      { x.swap(i, j) }

func laterrrr() {

	// sort them cookies
	cookies := CookieSlice{{10, "Chocolate", 5}, {12, "Peanuts", 4}, {8, "Almonds", 3}}
	sort.Sort(CookieBySizeSlice(cookies))

	// sort any slice by any order
	sort.Sort(&FuncSorter{
		func() int { return len(cookies) },
		func(i, j int) bool { return cookies[i].Rating < cookies[j].Rating },
		func(i, j int) { cookies[i], cookies[j] = cookies[j], cookies[i] },
	})

	// type assertions
	var quacker Quacker = &Duck{}
	if _, ok := quacker.(*Duck); ok {
		fmt.Println("is duck")
	}

	// when assigning to a single value
	// panics if the assertion fails
	_ = quacker.(*Duck)

	// type switches
	switch x := quacker.(type) {
	case *Duck:
		fmt.Printf("%v is duck\n", x)
		break
	default:
		fmt.Printf("%v is definitly no duck\n", x)
		break
	}

	takeNap := func() {
		time.Sleep(100 * time.Millisecond)
	}

	// functions invoked with
	// go are executed concurrently
	go takeNap()
	go takeNap()
	go takeNap()

	// goroutines communicate by
	// exchanging messages over channels
	channel := make(chan int, 3)

	producer := func() {
		fmt.Println("sending value 1")
		channel <- 1
	}

	// the receive operation is blocking
	consumer := func() {
		value := <-channel
		fmt.Printf("received value %v\n", value)
	}

	go producer()
	go consumer()
	time.Sleep(1 * time.Second)

	// the producer can signal it is done
	// by closing the channel
	producer = func() {
		fmt.Println("closing channel")
		close(channel)
	}

	consumer = func() {
		if _, ok := <-channel; !ok {
			fmt.Println("channel was closed")
		}
	}

	go producer()
	go consumer()
	time.Sleep(1 * time.Second)

	// a buffer size is defined
	// when the channel is created
	channel = make(chan int, 3)

	// if the buffer gets filled
	// back pressure is applied on the producer
	// by making the send operation blocking
	producer = func() {
		for i := 0; i < 10; i++ {
			start := time.Now()
			channel <- i
			fmt.Printf("waited %v sending value %v\n", time.Since(start), i)
		}
		close(channel)
	}

	// this loop construct automatically
	// breaks when the channel gets closed
	consumer = func() {
		for value := range channel {
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("consumed value %v\n", value)
		}
		fmt.Println("channel was closed")
	}

	go producer()
	go consumer()
	time.Sleep(2 * time.Second)

	// channels with a buffer size of zero
	// block the send and received operations
	// until a message is exchanged
	synchronizationChannel := make(chan int)

	worker1 := func() {
		fmt.Println("worker1 stuff...")
		time.Sleep(100 * time.Millisecond)
		fmt.Println("synchronizing")
		synchronizationChannel <- 1
		fmt.Println("more worker1 stuff...")
	}

	worker2 := func() {
		fmt.Println("worker2 stuff...")
		time.Sleep(200 * time.Millisecond)
		fmt.Println("synchronizing")
		_ = <-synchronizationChannel
		fmt.Println("more worker2 stuff...")
	}

	go worker1()
	go worker2()
	time.Sleep(1 * time.Second)

	// channel types can be used to
	// enforce the message directions
	var _ chan<- int = channel
	var _ <-chan int = channel

	// looping concurrently
	workItems := []int{1, 2, 3, 4, 5, 6, 8}

	for _, workItem := range workItems {
		go func(capturedWorkItem int) {
			fmt.Printf("work item %v done\n", capturedWorkItem)
		}(workItem)
	}

	time.Sleep(1 * time.Second)

	// with results accumulation
	results := make(chan int, len(workItems))

	for _, workItem := range workItems {
		go func(capturedWorkItem int) {
			fmt.Printf("work item %v done\n", capturedWorkItem)
			results <- capturedWorkItem
		}(workItem)
	}

	accumulation := 0
	for range workItems {
		accumulation += <-results
	}
	fmt.Printf("accumulation %v\n", accumulation)

	time.Sleep(1 * time.Second)
}
