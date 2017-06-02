package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Very nice shit to show the blank identifier

	res, _ := http.Get("http://valsot.com/")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	fmt.Printf("%s", page)
}
