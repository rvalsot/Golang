# Golang

This file are notes from the Tod McLeod's Udemy Course on Go: [Learn How To Code: Google's Go](https://www.udemy.com/learn-how-to-code/learn/v4/overview) also, additional research to improve my knowledge of the language.

Tod's McLeod contact:
* [Github Page](https://github.com/GoesToEleven)

## Installing Go

In Ubuntu, installation and Go Path setting:
``` shell
$ sudo apt-get update -y
$ sudo apt-get install -y golang

$ echo 'export GOROOT=/usr/local/go' >> ~/.bashrc
$ echo 'export GOPATH=/$HOME/go' >> ~/.bashrc
$ echo 'export PATH=$PATH:$GOROOT/bin:$GOPATH/bin' >> ~/.bashrc
```
While in Fedora, installation & path:
``` sh
$ sudo dnf update -y
$ sudo dnf install -y go

$ mkdir -p $HOME/work
$ echo 'export GOPATH=$HOME/work' >> $HOME/.bashrc
$ source $HOME/.bashrc
```

Run in terminal `go env` to check `$GOPATH` is correctly set.

__Go Workspace__

It is composed by any folder (preferred to be just one) containing:
`Workspace`
* `bin` :
* `pkg` :
* `src` :
    * Folder `project_a` repository
    * Folder `project_b` repository, ... etc.

This structure allows you easy package managing & _namespacing_

To add packages from the community, additional to our Standard Library, run in terminal:

``` sh
$ go get name_of_the_package
```

### Go Command Line

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

__Go Documentation__

[List of Documentation sites]

## Packages

To remember:
* Folder names containing the packages' files need to have the same name as the package.
* Package declaration must precede each file
* Upper case functions are exported, while lower-case are just available within the package.

## Variables

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

__Compound Types__

__Scope__

By ascendant order:
* Block: is available within curly braces, like inside a function. Within each block, precedence matters in being available.
* File
* Package: is done directly on _package level_ outside of a `main()` or any function.
* Universe

__Initialization__

There are a bunch of ways to initialize your variables:


## Functions

Anonymous functions or _func expressions_, assigned to a variable, without giving it a proper name.

``` go
package main

func main () {
  my_function := func() int {
    // Do something
    return x
  }
}
```

What can you return:
* Another function

## Creating a First Project

Suggested to do by McLeod:
* Create below `src/project_name/`
* Folder & Package name should correspond.
* Packages are given, by convention, lower-case and single word names.

## Resources

Extra web pages providing useful stuff.

* [Ardan Labs](https://www.ardanlabs.com/) $\leftarrow$ good challenges.
* [Fedora installation](https://developer.fedoraproject.org/tech/languages/go/go-installation.html)
* [Scope](https://golang.org/ref/spec#Declarations_and_scope)
* [Ubuntu installation](https://github.com/golang/go/wiki/Ubuntu)

## Miscellany
