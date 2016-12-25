package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Print("Hashing Moby Dick")
	fmt.Println(".")

	// Get the Book
	var theUrl string = "http://www.gutenberg.org/files/2701/old/moby10b.txt"
	res, err := http.Get(theUrl)

	// Look for errors
	if err != nil {
		fmt.Println("Aborting, fatal error")
		log.Fatal(err)
	}

	// Scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	// Set the split function for the scanning operation
	scanner.Split(bufio.ScanWords)

	// create slice to hold counts
	buckets := make([]int, 200)

	// Loop over the words
	for scanner.Scan() {
		n := HashBucket(scanner.Text())
		buckets[n]++
	}

	fmt.Println(buckets[65:122])

	fmt.Println("--------------------------------------------------")
	for i := 28; i <= 126; i++ {
		fmt.Printf("%v - %c - %v \n", i, i, buckets[i])
	}
}

func HashBucket1(word string) int {
	return int(word[0])
}
