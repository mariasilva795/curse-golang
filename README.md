Curse golang (https://www.udemy.com/certificate/UC-d2b9935c-2a1d-4d85-bef5-6ba70b14c14c/)

# Goland - **Deeper into Go**

### What is the Packages?

Every Go program is made up of packages. 

Programs start running in package `main`.

This program is using the packages with import paths `"fmt"` and `"math/rand"`.

### What is Golang?

- Powerfull standard library
- Concurrency management through  GoRoutines and Channels
- Designed by Google, Ken Thompson (UNIX) and part of design group
- Used in CLI and backend apps.
- Docker, Kubernetes y Terraform writter in Go.
- According to stackoverflow, Its the third highest paying language globally
- According to StackOverflow its fifth most loved tecnology  by  developers and the third most desired to work with.

## Defer

A defer statement defers the execution of a function until the surrounding function returns.

## Pointers

Go has pointers. A pointer holds the memory address of a value.

## Sección 5: **Organizing Data With Structs**

## Structs

A `struct` is a collection of fields.

## Arrays

The type `[n]T` is an array of `n` values of type `T`.

The expression

```
var a [10]int
```

declares a variable `a` as an array of ten integers.

**An array's length is part of its type, so arrays cannot be resized.** 

## Slices

When you define a slice variable in Go, such as `var mySlice []int`, Go creates a data structure that contains three pieces of information:

1. A pointer to an underlying array that holds the slice elements
2. A length that represents the number of elements in the slice
3. A capacity that represents the maximum number of elements that the underlying array can hold

For example, if you create a slice like this:

```
mySlice := []int{ 2, 3,4}
```

Go creates an underlying array like this:

```

+---------+---+
| 2| 3| 4|
+---------+--- +
  len=3, cap=4
```

The slice header contains a pointer to the start of the array (in this case, the second element), as well as the length and capacity of the slice.

## Creating a slice with make

Slices can be created with the built-in `make` function; this is how you create dynamically-sized arrays.

## Sección 5: Mapas

A "map" refers to a built-in data structure called a map. A map is a collection of key-value pairs, where each key must be unique. It provides an efficient way to look up values based on their associated keys.

## Sección 6: Interfaces

An *interface type* is defined as a set of method signatures.

A value of interface type can hold any value that implements those methods.

## Sección 7: Channels y Go Routines

Goroutines are a key feature of the Go programming language (Golang) designed to enable concurrent execution of functions. They are lightweight threads managed by the Go runtime, allowing developers to write concurrent code in a more straightforward and efficient manner.

Key characteristics of Goroutines include:

1. **Concurrency**: Goroutines allow developers to run functions concurrently, meaning multiple functions can be executed independently at the same time.
2. **Low Overhead**: Goroutines have lower overhead compared to traditional threads, making it practical to use a large number of them concurrently. This is facilitated by the Go scheduler, which efficiently multiplexes Goroutines onto a smaller number of operating system threads.
3. **Simplicity**: Launching a Goroutine is as simple as using the **`go`** keyword followed by the function call. This simplicity encourages developers to write concurrent code without the complexity often associated with threads and locks.

## Channels

For facilitating communication and synchronization between Goroutines (concurrently executing functions). Channels provide a way for one Goroutine to send data to another Goroutine in a safe and coordinated manner.
