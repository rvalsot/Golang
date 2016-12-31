# Golang

## Sources

This file are notes from the Tod McLeod's Udemy Course on Go: [Learn How To Code: Google's Go](https://www.udemy.com/learn-how-to-code/learn/v4/overview) also, additional research to improve my knowledge of the language.

Tod's McLeod contact:
* [Github Page](https://github.com/GoesToEleven)

## ✓ Installation and Workspace

__Installation__

Download from source with terminal and set the `$GOPATH` environment variable that will locate your workspace.

In Ubuntu:
``` shell
$ sudo apt-get update -y
$ sudo apt-get install -y golang

$ echo 'export GOROOT=/usr/local/go' >> ~/.bashrc
$ echo 'export GOPATH=/$HOME/go' >> ~/.bashrc
$ echo 'export PATH=$PATH:$GOROOT/bin:$GOPATH/bin' >> ~/.bashrc
```
While in Fedora:
``` sh
$ sudo dnf update -y
$ sudo dnf install -y go

$ mkdir -p $HOME/work
$ echo 'export GOPATH=$HOME/work' >> $HOME/.bashrc
$ source $HOME/.bashrc
```

Run in terminal `go env` to check `$GOPATH` is correctly set.

__Go Workspace__

It is composed by a folder at `$GOPATH`, containing sub directories:
` Workspace/`
* `/bin` : executable commands;
* `/pkg` : package objects;
* `/src` : Go source files;
    * Folder `project_a` repository
    * Folder `project_b` repository
    * Folder `package_a` repository, ... etc.

This structure allows you easy package managing & _namespacing_.

__Import Paths__

Import paths are unique identifiers for our packages origin. If having a Github account, it is recommended:
``` sh
$ mkdir -p $GOPATH/src/github.com/you/package_name
```

To add libre packages from the community, additional to our Standard Library, run in terminal:

``` sh
$ go get name_of_the_package
```


__Package Installation__

You can install code at your `/src` directory by using:
``` sh
$ go install /src/path_to_folder/your_package
```

Or by executing the `go install` command while in the directory.

After that, you can run the code at Terminal with:
``` sh
$ $GOPATH/bin/your_package
```

__Making a Library__

First, create your package stuff as normal Go files.

``` go
package your_package

// YourFunction does something ←this is your documentation
func YourFunction {
  // Do something
}
```

We build it at Terminal, if something is wrong, build will be halted and we notified.

``` sh
$ go build path_to_folder/your_package
```

After that, we can call our package using `import "path_to_folder/your_package"` in our scripts, as any other Standard Library.

Remember:
* Documentation if highly enforced as comments above each of our functions;
* Scoping is given by casing in the packages;

### Notes:
* `GOPATH` must not be the same location of your installation.


## Go Command Line

| Command | Short Description |
| ------- | ----------------- |
| `go build`  | Compiles packages and dependencies |
| `go clean`  | Remove object files (like older executables) |
| `go doc`    | Show the documentation for package of symbol |
| `go env`    | Print the Go Environment information |
| `go fix`    | Run go tool fix on packages |
| `go fmt`    | Formats the code according to Go's Standard |
| `go generate` | Generate Go files by processing the source |
| `go get`      | Download and install packages and dependencies |
| `go install`  | Compile and install packages and dependencies  |
| `go list`     | Lists your packages |
| `go run`      | Compile and Run a Go program |
| `go test`     | Test packages |
| `go tool`     | Runs a specific Go tool  |
| `go version`  | Prints your current Go version |
| `go vet`      | Run Go tool vet on packages |


__The Big Three for Go Terminal__
`go run`

Needs a file name as input, i.e.: `go run my_main.go`, will report errors, if they exists.

`go build`
1. For an _Executable_:
    * Builds the file, produces output file
    * Reports errors, if they exist, this will halt output
2. For a _Package_:
    * Builds the file, producing a binary
    * If any error, reports them and halts compilation.

It does not install the resulting stuff.
If the argument is a list of files, go treats them as a single package.

`go install`
1. For an _Executable_:
    * Compiles the program
    * Names the executable the folder name holding the code
    * Puts the executable in: `$GOPATH/bin`
2. For a _Package_:
    * Compiles the package
    * Puts the package in `$GOPATH/pkg`
    * Makes it an archive file

## Go Documentation

[List of Documentation sites]

## Packages

To remember:
* Folder names containing the packages' files need to have the same name as the package.
* Package declaration must precede each file
* Upper case functions are exported, while lower-case are just available within the package.

## Variables

✓ __Blank identifier__

Represented by the underscore character: `_`, is a special identifier you can assign anything you want, but never read from. It has some special uses in declarations, as and operand, and in assignments.

__Constants__

_A constant is a simple, unchanging value_, it is identified by the keyword `const` and can be declared individually or in block:

``` go
const p string = "This is a constant!"
const (
  a = 1992
  b = "Constant B"
)
```

The Iotas are auto-incremented assign values that can be used in constant blocks declaration with theeir keyword `iota`.
``` go
const (
  a = iota  // a = 0
  b = iota  // b = 1, and so on
)
```
Whenever `const` keyword appears, Iota count is reset to `0`.

__Primitive types__

| Group | Types |
| ----- | ----- |
| Boolean   | `bool` |
| Char      | `string` |
| Integers  | `int`, `int8`, `int16`, `int32`, `int64` |
| Unsigned Integers | `uint`, `uint8`, `uint16`, `uint32`, `uint64`; `uintptr` |
| Floating Point | `float32`, `float64` |
| Complex | `complex64`. `complex128` |
| Pointer |
| Others  |

__Runes__

They are the integer values identifying and Unicode code point, they're expressed in between single quotes.

__Compound Types__

*Or reference types?*

✓ __Scope__

By ascendant order:
* Block: is available within curly braces, like inside a function. Within each block, precedence matters in being available.
* File: the scope of a package name of an imported package is the file block of the file containing the import declaration.
* Package: is done directly on _package level_ outside of a `main()` or any function.
* Universe: the scope of a predeclared identifier is the _Universe block_.

__Initialization__

There are a bunch of ways to initialize your variables:

__Zero values__

### Data Structures (Reference type)

Are constructs of an underlying data type

* Array
* Slice
* Map
* Struct

__Common Stuff__

* `append(a, b, c...)`, will paste the elements, you can paste slices, for example
  * Appending the same slice, but skipping a given index will delete it.
* `cap(a)`, will show you the capacity of the structure
* `len(a)`, will give you the
* `make(a, n, m)`,  can be used to create channels, maps, and slices
* Length / Capacity difference:

__Operations__

__Arrays__

The length of the array is part of its type.
They're linear, but can be composed of multidimensional stuff.

``` go
myArray := [size]type
```

__Slices__

[# Definition]

Declaration and access

``` go

// Declaration
var mySlice []type                  // var style
mySlice := []type{}                 // short style
mySlice := make([]type], len, cap)  // make style

multidimensionalSlice := [][]type{}         // short style
multidimensionalSlice := make([][]type, 0)  // make style

// Filling
mySlice = append(mySlice, someStuff)
mySlice[i] = someStuff

// Access to elements
mySlice[n]                  // will access n + 1 element
multidimensionalSlice[n]    // will access the n + 1 dimension
multidimensionalSlice[n][m] // will access the n +1, m + 1 element
```

To remember:

* A slice done using `make` will duplicate its capacity if overflown.
* Just when using `make`, we can assign by index, if not, gotta use `append` to add stuff to our slices.
* Using `make(name, n)`, will give it the same length and capacity.
* You can increment by index: `mySlice[i]++` an element.

__Maps__

Are Key-Value storages; an unordered group of elements of one type, called _element type_, indexed by a set of _unique_ keys of another type, called _key type_. Their uninitialized value is `nil`.

``` go
var myMap = make(map[key_type]value_type)     //
var myMap = map[key_type]value_type{
  key:value,
  key2:value2,
}                                            //

```
Checking existance
```
[c96 exists if ]
```

To remember:
* You can have maps of maps.
* Using `make` to create a map, you can set an approximate length

__Hash Tables__

[Gotta Put something here]

__Struct__

Are user-defined, composite types, created this way:
``` go
type name struct {
  field   type
  field2  type // ... and so on...
}
```
Field names must be unique and be explicitly or implicitly defined.

``` go
variable := myStruct{1,2,"a",true, inherited_struct}      // Short declaration
variable := myStruct{
  field_i : 1,
  field_i : 2,
  field_s : "a",
  field_b : true,
  inherited_struct : inherited_struct {
    field_o : "o"
  }
}
```
To know:
* Structs get subjected to the same Export / Not behavior than functions.
* You can attach methods to a type.
* There's inheritance of fields.
* Remember with inheritance: _The Outer type overwrites the Inner type_

Inherited stuff:
``` go
type provider {
  field type // ... and so on
}

type receiver {
  provider
  field type // and so on
}
newProvider := prodiver{quality}
newReceiver := receiver{newProvider, quality2} // or
newReceiver := receiver{provider:provider{}, quality2}
```

Tags:

__Struct Methods__

Are declared a almost normal functions:

``` go
func (a struct) MethodName() type {
  // do something
}
```
Any value of the type `struct` can use the `MethodName` as proper.

__Interfaces__

They are the Go response to polymorphism: the ability to write code that can take on different behavior through the implementation of types.

Interfaces are types that just declare behavior. They're implemented by user-defined types via Methods.

They have not

[Empty Interfaces <- Everything implements them]
(Pointer receivers)

[Method Sets]

## Memory

__Pointers__

``` go
a := 2
var b *int = &a
```
Here, the pointer `b` is directed to the integer `a` address'.

## Channels



## Flow Control

__For Loop__

The `for` loop in Go has several behaviors, accomplishes for, while, for-each and do-while like actions from other languages:
``` go
for init; condition; post {
  // this is a common C-like for
}

for i < 10 {
  i++ // This will be the equivalent of while in C languages, just need a Boolean condition!
}

for key, value := range aMap {
  // This is a for-each, to iterate in an: array, slice, string or map, or reading a channel
}

for key := range anArray {
  // this is similar to previous for, but will only iterate the first element of the range
}

for _, value := range aRange {
  // Similar to both previous, will drop the first element and go for the second, with the blank identifier
}

for {
  // This for will run indefinitely, like the C for(;;)
  continue // This will begin the next iteration of the loop
  break // This statement can halt the execution and exit the loop
}
```

__Switch__

There are several versions of the `switch` statement, the standard one is:
``` go
switch expression {
  case condition:
    // Do something
  default: the_default
    // Do else
}
```
There's also an empty switch that will execute the commands that first (and just this) return `true` value:
``` go
switch {
  case false:
    // This will not execute
  case true:
    // This will execute
  default:
    // If previous were false, this will run   
}
```

__If blocs__

We have three forms:
* Simple `if`
* Compound `if`, `else`
* Compound `if`, `else if`, `else`

``` go
if condition {
  // Do something, this is a simple if, will execute if "condition" is true
}

if initialize; condition {
  // Will initialize something before the expression is evaluated.
}
```

The composite versions would be:
``` go
if condition {
  // Do a-stuff
} else if condition {
  // Do b-stuff
} else {
  // do the last-resort stuff
}

```

__Defer Statement__

The `defer` statement preceding a function, halts its execution to be ran immediately before the function execution the defer returns; if there are several defer, they will execute in regressive order:
``` go
func main () {
    defer fmt.Println("1")
    defer fmt.Println("2")
    fmt.Println("3")
}
// This will print out:
3
2
1
```

__Boolean Expressions__

* Logical Not `!`
* Logical And `&&`
* Logical Or  `||`

## Functions

| Receiver | Values |
| -------- | ------ |
| `(t T)`  | `T` and `*T` |
| `(t *T)` | `*T` |

| Values | Receivers |
| ------ | --------- |
| `T`    | `(t T)`   |
| `*T`   | `(t T)` and `(t *T)` |



__Anonymous functions__
Also called _func expressions_, assigned to a variable, without giving it a proper name.

``` go
package main

func main () {
  my_function := func() int {
    // Do something
    return x
  }
}
```
They are used whenever we want to have a function declaration inside another, like inside our `main()`.

Self executing, Anonymous Functions are not declared nor named:
``` go
func () {
  // Do something
} () // This parenthesis will run the function
```

__Function Expressions__

Are functions that can be assigned to a variable and whose return type is another function and its own type:

``` go
function myFunction() func int(){
  return func () int {
    x++
    return x
  }
}

myVariable := myFunction()
```

__Functions Proper__

The basic structure of a function in go is:
``` go
func function_name (parameter type, param2 type) (returnType, otherReturnType) {
  // does something
  return something
}
```
Where:
* First, we declare the function name
* After that, we list the input parameters followed by their type and separated by commas, a series of parameters of the same type can be listed and just one type added at the end.
* The third part is the return: we can name it or not, but we need to specify the return type always; if there are more than one returned value, we must list them between parenthesis.

__Variadic Functions__

Are functions that can take from zero to an unspecified number or arguments of the same type, they are declared this way:
``` go
func myFunction (name ...type) returnType {
  // Does some stuff
}
```

Remember that passing a group of arguments it's not the same that passing a slice!

__Callback functions__

Callbacks allow us to pass functions as an argument:
``` go
func baseFunction(parameters type, callback func(type)) baseFunctionReturnType {
  callback(argument)
  return something
}
// We use it, lets say:
baseFunction(our_arguments, func(callback_payload) {
  fmt.Println(callback_payload)
})

```

__Pass by value__

Sending the address of a variable

Strings are immutable.

__Return Values__
What can you return:
* Another function

Named returns do not need to `return` keyword to be followed by the variable, just the last being assigned somewhere in the function body:
``` go
func aFunction (parm type) (returnedValue type) {
  returnedValue := something //
  return
}
```

You can assign the return value to a variable in the context is being used:
``` go
myVariable := myFunction(arguments)
```

__Recursion__

Go supports Recursion, simply put it in the return!
``` go
func factorial(x int) int {
  if x == 0 {
    return 1
  }
  return x * factorial(x-1)
}
```

__Notes__

* Everything in Go is passed by copy.


## Concurrency

> Concurrency is the composition of independently executing processes, while parallelism is the simultaneous execution of, possibly related, computations. First is about dealing with lots of things at once; Second about doing lots of things at once.

To split our programs and use concurrency, Go provides us of Goroutines and Channels.

__Goroutines__
Are capable of running concurrently with other functions, they're created with the keyword `go`.

A

To avoid Concurrency overwriting

[WaitGroups]
[Atomicity -> sync package]
[Mutex -> mutex package]

## Resources and Miscellany

* [Ardan Labs](https://www.ardanlabs.com/) ← good challenges.
* [Fedora installation](https://developer.fedoraproject.org/tech/languages/go/go-installation.html)
* [Going Go](https://www.goinggo.net/)
* [Golang Book](https://www.golang-book.com/books/intro) ← The gopher in the bike.
* [Scope](https://golang.org/ref/spec#Declarations_and_scope)
* [Ubuntu installation](https://github.com/golang/go/wiki/Ubuntu)
