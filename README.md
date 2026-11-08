# Go: Function Anatomy & Anonymous Functions

A comprehensive guide to understanding function signatures, parameter grouping, multiple return values, and anonymous functions in Go (Golang).

---

## 📑 Table of Contents

- [Overview](#overview)
- [The Anatomy of a Function](#the-anatomy-of-a-function)
- [Key Concepts & Examples](#key-concepts--examples)
  - [1. Standard Function Declaration](#1-standard-function-declaration)
  - [2. Parameter Grouping (Type Omission)](#2-parameter-grouping-type-omission)
  - [3. Multiple Return Values](#3-multiple-return-values)
  - [4. The Blank Identifier (`_`)](#4-the-blank-identifier-_)
  - [5. Named Return Values](#5-named-return-values)
  - [6. Anonymous Functions & Closures](#6-anonymous-functions--closures)
- [Code Walkthrough (`main.go`)](#code-walkthrough-maingo)
- [How to Run](#how-to-run)
- [Quick Reference Cheat Sheet](#quick-reference-cheat-sheet)

---

## Overview

Functions are the primary building blocks in Go. They allow you to structure code into reusable, modular, and maintainable units. Unlike many traditional languages, Go treats functions as first-class citizens and natively supports features such as **multiple return values**, **parameter type grouping**, and **anonymous functions (closures)**.

---

## The Anatomy of a Function

Every function in Go follows a clear and consistent signature:

```go
func FunctionName(param1 Type, param2 Type) (ReturnType1, ReturnType2) {
    // Function body
    return val1, val2
}
```

### Breakdown:

| Component | Description |
| :--- | :--- |
| `func` | The keyword used to declare a function. |
| `FunctionName` | The identifier used to call the function (starts with lowercase for package-private, uppercase for exported). |
| `Parameters` | Values passed into the function along with their types enclosed in `(...)`. |
| `Return Types` | The type(s) of values the function returns. If returning multiple values, wrap them in parentheses `(...)`. |
| `{ ... }` | The code block (body) that executes when the function is invoked. |

---

## Key Concepts & Examples

### 1. Standard Function Declaration

Parameters in Go are declared with the parameter name followed by its type:

```go
func test(a int, b string, c int, d int) string {
    return fmt.Sprintf("%d %s %d %d", a, b, c, d)
}
```

- Each parameter has an explicit type declared right after the parameter name.
- Single return value: Does not require parentheses around the return type.

---

### 2. Parameter Grouping (Type Omission)

When consecutive parameters share the same type, you can omit the type for all but the last parameter:

```go
// Instead of: func test3(a int, b int, c int, d int) int
func test3(a, b, c, d int) int {
    return a + b + c - d
}
```

> **Tip:** This keeps signatures clean, concise, and easy to read.

---

### 3. Multiple Return Values

Go natively supports returning multiple values from a single function. This is commonly used in Go for returning results alongside errors:

```go
func tren(a int, c int) (int, int) {
    sum := a + c
    multion := a * c
    return sum, multion
}
```

You unpack multiple return values using comma-separated variable declarations:

```go
sum, multion := tren(23, 56)
```

You can return values of different types as well:

```go
func tren4(a int, c int, d float64, b int) (int, int, float64, int) {
    sum := a + c
    multion := a * c
    ma := float64(int(d) % b)
    re := c / b

    return sum, multion, ma, re
}
```

---

### 4. The Blank Identifier (`_`)

If a function returns multiple values and you only need some of them, use the blank identifier `_` to discard the ones you don't need:

```go
// Only care about the sum, discard the multiplication result
sum, _ := tren(10, 20)
```

---

### 5. Named Return Values

Go allows you to name the return values in the function signature. They are treated as variables defined at the top of the function:

```go
func calculate(a, b int) (sum int, product int) {
    sum = a + b
    product = a * b
    return // Naked return: returns sum and product automatically
}
```

---

### 6. Anonymous Functions & Closures

An **anonymous function** (or function literal) is a function defined without an identifier/name.

#### A. Assigned to a Variable
You can assign an anonymous function to a variable and call it later:

```go
greet := func(name string) string {
    return "Hello, " + name
}

fmt.Println(greet("Developer")) // Output: Hello, Developer
```

#### B. Immediately Invoked Function Expression (IIFE)
You can define and execute an anonymous function on the spot by placing `()` at the end:

```go
func(message string) {
    fmt.Println(message)
}("Executing immediately!")
```

#### C. Closures
Anonymous functions can access and modify variables declared outside their body:

```go
counter := 0
increment := func() int {
    counter++
    return counter
}

fmt.Println(increment()) // 1
fmt.Println(increment()) // 2
```

---

## Code Walkthrough (`main.go`)

Here is the source code breakdown of `main.go`:

```go
package main

import (
	"fmt"
)

// Demonstrates mixed parameter types and string formatting
func test(a int, b string, c int, d int) string {
	return fmt.Sprintf("%d %s %d %d", a, b, c, d)
}

// Demonstrates grouped parameter types (a, b, c, d are all ints)
func test3(a, b, c, d int) int {
	return a + b + c - d
}

// Demonstrates multiple return values of the same type (int, int)
func tren(a int, c int) (int, int) {
	sum := a + c
	multion := a * c

	return sum, multion
}

// Demonstrates multiple return values of different types (int, int, float64, int)
func tren4(a int, c int, d float64, b int) (int, int, float64, int) {
	sum := a + c
	multion := a * c
	ma := float64(int(d) % b)
	re := c / b

	return sum, multion, ma, re
}

func main() {
	fmt.Println("Hello")

	// Calling a single-return function
	point := test(23, "wegdfy", 777, 5)
	fmt.Println(point)

	// Unpacking multiple return values
	sum, multion := tren(23, 56)

	// Re-assigning / declaring new return values
	sum, multion, ma, ra := tren4(23, 56, 90.877, 676)
	fmt.Println("Sum:", sum)
	fmt.Println("Multiplication:", multion)
	fmt.Println("Modulus:", ma)
	fmt.Println("Division:", ra)

	fmt.Println(sum, multion)

	// Example of an Anonymous Function:
	func() {
		fmt.Println("Anonymous function executed!")
	}()
}
```

---

## How to Run

1. Make sure you have [Go](https://go.dev/dl/) installed:
   ```bash
   go version
   ```

2. Run the file directly:
   ```bash
   go run main.go
   ```

3. (Optional) Build an executable binary:
   ```bash
   go build -o app
   ./app
   ```

### Expected Output

```text
Hello
23 wegdfy 777 5
Sum: 79
Multiplication: 1288
Modulus: 90
Division: 0
79 1288
```

---

## Quick Reference Cheat Sheet

| Feature | Syntax Example |
| :--- | :--- |
| **Basic Function** | `func add(x int, y int) int` |
| **Type Omission** | `func add(x, y int) int` |
| **Multiple Returns** | `func divide(a, b int) (int, error)` |
| **Named Returns** | `func calc(a, b int) (sum int)` |
| **Ignore Value** | `val, _ := multiReturnFunc()` |
| **Anonymous Function** | `fn := func() { ... }` |
| **Self-Executing (IIFE)** | `func() { ... }()` |

---

## 📚 Author

- **GitHub:** [@aftabfarhanarko](https://github.com/aftabfarhanarko)
