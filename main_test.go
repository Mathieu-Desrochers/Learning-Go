package main

import (
	"fmt"
	"math"
	"testing"
)

// running tests
// go test
func TestAddition(t *testing.T) {
	if 2+2 != 4 {
		t.Error("2+2 != 4")
	}
}

func TestTableDriven(t *testing.T) {
	var tests = []struct {
		input float64
		want  float64
	}{
		{1, 1},
		{2, 4},
		{3, 9},
	}
	for _, test := range tests {
		if got := math.Pow(test.input, 2); got != test.want {
			t.Errorf("math.Sqrt(%v) = %v, want %v", test.input, got, test.want)
		}
	}
}

// mocking using global variables
var selectCustomer = func(customerId int) string {
	return fmt.Sprintf("SELECT Name FROM Customers WHERE Id = %v;", customerId)
}

func GetCustomer(customerId int) string {
	return "His name is " + selectCustomer(customerId)
}

func TestGetCustomer(t *testing.T) {
	selectCustomerReal := selectCustomer
	defer func() { selectCustomer = selectCustomerReal }()

	selectCustomer = func(customerId int) string { return "Bob" }

	got := GetCustomer(1)
	fmt.Printf("got %v\n", got)
}

// computing test coverage
// and displaying the green/red source
// go test -coverprofile=cover.out
// go tool cover -html=cover.out

// benchmarking time and memory allocations
// go test -bench=.
// go test -bench=. -benchmem

func Lengthy() {
	for i := 0; i < 10000000; i++ {
	}
}

func BenchmarkLengthy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Lengthy()
	}
}

// profiling CPU, memory and blocking
// go test -bench=. -cpuprofile=cpu.out
// go test -bench=. -memprofile=mem.out
// go test -bench=. -blockprofile=block.out
// go tool pprof -text -nodecount=10 ./cpu.out

// providing examples
// included in the documentation
// output checked when tests are run
func Division(x, y int) int {
	return x / y
}

func ExampleDivision() {
	fmt.Println(Division(4, 2))
	fmt.Println(Division(10, 2))
	// Output:
	// 2
	// 5
}
