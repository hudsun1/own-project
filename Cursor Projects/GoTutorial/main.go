package main

import (
	"fmt"
	"time"
)

// Basic Hello World example
func helloWorld() {
	fmt.Println("\n=== Hello World Example ===")
	fmt.Println("Hello, Go World!")
	name := "Go Developer"
	fmt.Printf("Welcome, %s!\n", name)
	a, b := 10, 5
	fmt.Printf("Sum: %d, Difference: %d, Product: %d\n",
		a+b, a-b, a*b)
}

// Variables example
func variablesDemo() {
	fmt.Println("\n=== Variables Example ===")
	var name string = "John"
	age := 25
	var isStudent bool = true

	var (
		city    string = "New York"
		country string = "USA"
	)

	const PI = 3.14159
	const (
		MAX_SIZE = 100
		MIN_SIZE = 1
	)

	fmt.Printf("Name: %s, Age: %d, Is Student: %t\n", name, age, isStudent)
	fmt.Printf("Location: %s, %s\n", city, country)
	fmt.Printf("PI: %.5f, Max Size: %d, Min Size: %d\n", PI, MAX_SIZE, MIN_SIZE)
}

// Control structures example
func controlStructuresDemo() {
	fmt.Println("\n=== Control Structures Example ===")

	// If-else
	age := 18
	if age >= 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a minor")
	}

	// Switch
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("Start of the week")
	case "Friday":
		fmt.Println("End of the week")
	default:
		fmt.Println("Middle of the week")
	}

	// For loop
	fmt.Println("\nCounting from 1 to 5:")
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

// Functions example
func functionsDemo() {
	fmt.Println("\n=== Functions Example ===")

	// Basic function
	result := add(5, 3)
	fmt.Printf("5 + 3 = %d\n", result)

	// Multiple return values
	if quotient, err := divide(10, 2); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", quotient)
	}

	// Variadic function
	total := sum(1, 2, 3, 4, 5)
	fmt.Printf("Sum of 1,2,3,4,5 = %d\n", total)
}

// Helper functions for functionsDemo
func add(x, y int) int {
	return x + y
}

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return x / y, nil
}

func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Structs and interfaces example
func structsAndInterfacesDemo() {
	fmt.Println("\n=== Structs and Interfaces Example ===")

	// Create a circle
	circle := Circle{Radius: 5}
	fmt.Printf("Circle with radius 5:\n")
	fmt.Printf("Area: %.2f\n", circle.Area())
	fmt.Printf("Perimeter: %.2f\n", circle.Perimeter())

	// Create a rectangle
	rect := Rectangle{Width: 4, Height: 6}
	fmt.Printf("\nRectangle 4x6:\n")
	fmt.Printf("Area: %.2f\n", rect.Area())
	fmt.Printf("Perimeter: %.2f\n", rect.Perimeter())
}

// Concurrency example
func concurrencyDemo() {
	fmt.Println("\n=== Concurrency Example ===")

	// Basic goroutine
	go func() {
		fmt.Println("Hello from goroutine!")
	}()
	time.Sleep(100 * time.Millisecond)

	// Channel example
	ch := make(chan string)
	go func() {
		ch <- "Hello from channel!"
	}()
	msg := <-ch
	fmt.Println(msg)
}

// Shape interface and implementations
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func main() {
	fmt.Println("Go Programming Language Tutorial")
	fmt.Println("===============================")

	// Run all examples
	helloWorld()
	variablesDemo()
	controlStructuresDemo()
	functionsDemo()
	structsAndInterfacesDemo()
	concurrencyDemo()

	fmt.Println("\nTutorial completed!")
}
