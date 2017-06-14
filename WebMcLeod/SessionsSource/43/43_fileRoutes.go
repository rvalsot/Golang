package main()

import (
  "fmt"
  "log"
  "net/http"
)

func main(){
  port := ":8070"
  fmt.Println("Listening & Serving at port", port)
  log.Fatal(
    http.ListenAndServe("",
      http.FileServer(
        http.Dir("."))))
}
