package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Niara the Hedgehog"
	message := "OS Templates"

	htmlString := fmt.Sprint(`
    <!DOCTYPE html>
    <html lang="en">
      <head>
        <meta charset="UTF-8">
        <title>Hi folks!</title>
      </head>

      <body>
        <h1>` + name + `</h1>
        <p> ` + message + `</p>
      </body>
    </html>
  `)

	newFile, err := os.Create("01_index.html")
	if err != nil {
		log.Fatal("Error creating the file → \n", err)
	}
	defer newFile.Close()
	io.Copy(newFile, strings.NewReader(htmlString))

}
