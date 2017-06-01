package main

import "fmt"

func main() {
	name := "Niara the Hedgehog"
	message := "01 templates"

	tpl := `
  <!DOCTYPE html>
  <html lang="en"/>
  <head>
    <meta charset="UTF-8"/>
    <title> Our first app! </title>
  </head>

  <body>
    <h1>` + name + `</h1>
    <p> ` + message + `</p>
  </body>
  </html>
  `
	fmt.Println(tpl)
}
