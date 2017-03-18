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
* `flag`:
  * `Parse`
  * `String`:
* `log`:
  * `Println`:
* `net/http`:
  * `HandleFunc`:
  * `ListenAndServe`:
  * `ResponseWriter`:
  * `Request`:
* `path/filepath`:
  * `Join`:
* `sync`:
  * `Once`:
* `template`
  * `ParseFiles`:
  * `Template`:
    * `Execute`:
* `delete(map[typeA]typeB, typeA)`:
* `close(chan)`:


## Other Libraries

* Gorilla
  * `socket`:
    * `Close`:
    * `ReadMessage`:
    * `Upgrader`:
    * `WriteMessage`:
