# Go Programming Language Tutorial

## Introduction to Go
Go (or Golang) is a statically typed, compiled programming language designed at Google. It's known for its simplicity, efficiency, and strong support for concurrent programming.

## Table of Contents
1. [Getting Started](#getting-started)
2. [Basic Syntax](#basic-syntax)
3. [Variables and Data Types](#variables-and-data-types)
4. [Control Structures](#control-structures)
5. [Functions](#functions)
6. [Structs and Interfaces](#structs-and-interfaces)
7. [Concurrency](#concurrency)
8. [Error Handling](#error-handling)

## Getting Started

### Installation
1. Download Go from [golang.org](https://golang.org/dl/)
2. Follow the installation instructions for your operating system
3. Verify installation by running:
```bash
go version
```

### Your First Go Program
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

## Basic Syntax

### Package Declaration
Every Go file starts with a package declaration:
```go
package main
```

### Imports
Import packages using the `import` keyword:
```go
import (
    "fmt"
    "math"
)
```

## Variables and Data Types

### Variable Declaration
```go
// Explicit declaration
var name string = "John"

// Type inference
age := 25

// Multiple variables
var (
    firstName string
    lastName  string
    age       int
)
```

### Basic Data Types
- `bool`: true/false
- `string`: text
- `int`, `int8`, `int16`, `int32`, `int64`: integers
- `uint`, `uint8`, `uint16`, `uint32`, `uint64`: unsigned integers
- `float32`, `float64`: floating-point numbers
- `complex64`, `complex128`: complex numbers

## Control Structures

### If Statements
```go
if age >= 18 {
    fmt.Println("Adult")
} else {
    fmt.Println("Minor")
}
```

### Switch Statements
```go
switch day {
case "Monday":
    fmt.Println("Start of week")
case "Friday":
    fmt.Println("End of week")
default:
    fmt.Println("Middle of week")
}
```

### Loops
```go
// For loop
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// While loop (using for)
count := 0
for count < 5 {
    fmt.Println(count)
    count++
}
```

## Functions

### Function Declaration
```go
func add(x, y int) int {
    return x + y
}

// Multiple return values
func divide(x, y float64) (float64, error) {
    if y == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return x / y, nil
}
```

## Structs and Interfaces

### Structs
```go
type Person struct {
    Name    string
    Age     int
    Address string
}

// Creating a struct instance
person := Person{
    Name:    "John",
    Age:     30,
    Address: "123 Main St",
}
```

### Interfaces
```go
type Shape interface {
    Area() float64
    Perimeter() float64
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.Radius
}
```

## Concurrency

### Goroutines
```go
func main() {
    go sayHello() // Runs concurrently
    time.Sleep(time.Second)
}

func sayHello() {
    fmt.Println("Hello from goroutine!")
}
```

### Channels
```go
func main() {
    ch := make(chan string)
    
    go func() {
        ch <- "Hello from channel!"
    }()
    
    msg := <-ch
    fmt.Println(msg)
}
```

## Error Handling

### Basic Error Handling
```go
func readFile(filename string) (string, error) {
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        return "", fmt.Errorf("failed to read file: %v", err)
    }
    return string(data), nil
}
```

## Best Practices
1. Use meaningful variable names
2. Follow Go's formatting conventions (use `gofmt`)
3. Write unit tests for your code
4. Handle errors explicitly
5. Use interfaces for abstraction
6. Keep functions small and focused
7. Use goroutines and channels for concurrency

## Resources
- [Official Go Documentation](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [Go Tour](https://tour.golang.org/)
- [Effective Go](https://golang.org/doc/effective_go)

## Next Steps
1. Practice writing small programs
2. Learn about Go modules and dependency management
3. Explore the standard library
4. Study design patterns in Go
5. Contribute to open-source Go projects 