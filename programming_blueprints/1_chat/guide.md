# Chapter 1

Chat Application with Web Sockets

## Objectives

* Use the `net/http` package to serve HTTP requests
* Deliver template-driven content to user's browsers
* Satisfy a Go interface to build our own `http.Handler` types
* Use Go's `go routines` to allow an application to perform multiple tasks concurrently
* Use channels to share information between running Go routines
* Upgrade HTTP requests to use features such as web sockets
* Add tracing to the application to better understand its inner workings
* Write a complete Go package using test-driven development practices
* Return un exported types through exported interfaces

Credits goes to: https://github.com/matryer/goblueprints/tree/master/chapter1/chat

## Brotips, notes, etc.

* __Channels__: you can think them as in-memory and thread-safe message queue, where users pass and receive data  in a non-blocking thread-safe way.

## Used functions of the Standard Library
* `bytes`: manipulations for byte slices. Analogous to the `strings` package.
  * `Buffer`: is a variable-sized buffer of bytes with Read & write methods.
* `flag`: command-line flags parsing.
  * `Parse`: parses the command-line flags from OS. Must be called after all flags are defined and before flags are accessed by the program.
  * `String`: defines a string flag with specified name, default value, and usage string. Its return is the address of a string variable that stores the value of the flag.
* `io`: provides basic interfaces to IO primitives to wrap existing implementations.
  * `Writer`: is an interface that wraps the basic Write method. Writes $len(p)$ byte from `p` to the underlying data stream. And returns the number of bytes written and any error that may cause an early stop.
* `log`: implements a simple logging service.
  * `Println`: Calls Output to print to the standard logger. Equivalent to `fmt.Println()`.
* `net/http`: provides HTTP client and server implementation.
  * `HandleFunc`:
  * `Handler`:
  * `ListenAndServe`:
  * `ResponseWriter`:
  * `Request`:
* `path/filepath`: implements utility routines for manipulating filename paths in a way compatible with the target operating system-defined pile paths.
  * `Join`: joins any number of path elements into a single path, adding Separators if necessary.
* `sync`: basic synchronization primitives such as mutual exclusions locks.
  * `Once`: this type performs exactly one action. `t.Once.Do(f func(){})` calls the function `f` if and only if `Do` is being called for the first time for this instance of Once. Is intended for initialization that must be run exactly once.
* `text/template`: implements data-driven templates for generating textual output. `html/template` has the same interface, but secures HTML output against certain attacks.
  * `ParseFiles`:
  * `Template`:
    * `Execute`:
* `testing`: support for automated testing of Go packages.
  * `T`: is a type passed to the Tests functions to manage test test state and support formatted test logs, which are accumulated during execution and dumped to standard output when done.
* `delete(map[typeA]typeB, typeA)`:
* `close(chan)`:


## Other Libraries

* Gorilla
  * `socket`:
    * `Close`:
    * `ReadMessage`:
    * `Upgrader`:
    * `WriteMessage`:
