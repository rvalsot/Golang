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


* `net/http`
  * `HandleFunc`
  * `ListenAndServe`
* `path/filepath`
  * `Join`
* `sync`
  * `Once`
* `template`
  * `ParseFiles`
  * `Template`
    * `Execute`
