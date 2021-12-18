package main

import (
	"bytes"
	"errors"
	"fmt"
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
	fmt.Printf("an array of %v elements\n", len(array))

	// array literals
	_ = [3]int{1, 2, 3}
	_ = [...]int{1, 2, 3, 4}
	_ = [...]int{2: 10, 4: 20}

	// arrays are compared by value
	fmt.Printf("arrays a equal: %v\n", [...]int{1, 2} == [...]int{1, 2})

	// and are passed by value
	originalArray := [...]string{"a", "b", "c"}
	copiedArray := originalArray
	originalArray[0] = "x"
	fmt.Printf("originalArray: %v\n", originalArray)
	fmt.Printf("copiedArray: %v\n", copiedArray)

	// slices are pointers to an array
	// they keep track of it's length and capacity
	var slice []int
	fmt.Printf("a slice of %v elements and a capacity for %v\n", len(slice), cap(slice))

	// slice literals
	_ = []int{}
	_ = []int{1, 2, 3, 4}
	_ = []int{2: 10, 4: 20}

	// appending values
	// can reallocate the array to a different size and location
	// must be recaptured
	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3)
	fmt.Printf("appended slice %v\n", slice)

	// selecting values
	fmt.Printf("selected slice %v\n", originalArray[1:])

	// selected slices can get disconnected
	// if their source is reallocated
	sourceSlice := []int{1}
	selectedSlice := sourceSlice[0:1]
	sourceSlice = append(sourceSlice, 11)
	sourceSlice[0] = 10
	fmt.Printf("source slice %v\n", sourceSlice)
	fmt.Printf("selected slice %v\n", selectedSlice)

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
	fmt.Printf("a slice of %v elements and a capacity for %v\n", len(slice), cap(slice))

	// maps are hash tables
	var nameById = make(map[int]string)

	// map literals
	_ = map[int]string{}
	_ = map[int]string{1: "Alice", 2: "Bob"}

	// setting a value
	nameById[100] = "Alice"
	nameById[200] = "Bob"
	nameById[300] = "Carl"

	// looking up a value
	if name, ok := nameById[100]; ok {
		fmt.Printf("name: %v\n", name)
	}

	// iterating over values
	// order is not guaranteed
	for id, name := range nameById {
		fmt.Printf("id: %v, name: %v\n", id, name)
	}

	// removing a value
	delete(nameById, 300)

	// strings are immutable sequences of bytes
	greek := "Some greek: Τη γλώσσα μου έδωσαν"

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
	originalEmployee := Employee{FirstName: "A"}
	copiedEmployee := originalEmployee
	originalEmployee.FirstName = "X"
	fmt.Printf("originalEmployee: %v\n", originalEmployee)
	fmt.Printf("copiedEmployee: %v\n", copiedEmployee)

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

// declared errors
var outOfPaint = errors.New("out of paint")

func ooopsAgain() error {
	return outOfPaint
}
func errorDeclared() {
	err := ooopsAgain()
	if err == outOfPaint {
		fmt.Println(err)
	}
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

	// errors
	errorWithContext("red")
	errorDeclared()

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
		defer fmt.Println("deferred")
		{
			defer fmt.Println("does not work at block level")
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
func (a Animal) CanQuack() bool {
	return false
}

// the receiver is passed by value
// methods are primarily used with pointers
func (a *Animal) GrowLeg() {
	a.LegsCount++
}

func laterr() {

	// methods
	animal := &Animal{4}
	fmt.Println(animal.CanQuack())

	animal.GrowLeg()
	fmt.Println(animal.LegsCount)

}
