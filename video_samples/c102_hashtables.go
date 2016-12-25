package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Print("Hash Tables")
	fmt.Println("")

	theUrl := "https://raw.githubusercontent.com/Pr0x13/iDict/master/files/wordlist.txt"

	res, err := http.Get(theUrl)
	if err != nil {
		fmt.Println("Error")
		log.Fatalln(err)
	}

	bs, _ := ioutil.ReadAll(res.Body)
	str := string(bs)
	// fmt.Print(str)
	fmt.Printf("%T \n", str)

	words := make(map[string]string)
	sc := bufio.NewScanner(res.Body)
	sc.Split(bufio.ScanWords)

	for sc.Scan() {
		words[sc.Text()] = ""
	}

	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input", err)
	} else {
		fmt.Println("Was read.")
	}
}
