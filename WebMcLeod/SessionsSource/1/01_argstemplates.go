package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := os.Args[1]
	message := os.Args[2]
	fmt.Println(os.Args[0], "← 0")
	fmt.Println(os.Args[1], "← 1")

	htmlString := fmt.Sprint(`
    <!DOCTYPE html>
      <html lang="en">
      <head>
        <meta charset="UTF-8">
        <title>Hello World!</title>
      </head>
      <body>
        <h1>` + name + `</h1>
        <p> ` + message + `</p>
      </body>
    </html>
  `)

	newFile, err := os.Create("01_index.html")
	if err != nil {
		log.Fatal("Error creating the file: \n", err)
	}
	defer newFile.Close()
	io.Copy(newFile, strings.NewReader(htmlString))

}
