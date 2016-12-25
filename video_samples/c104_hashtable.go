package main

import "fmt"

func main() {
	fmt.Print("Hash Table")
	fmt.Println("")
	n := HashBucket("Golang", 12)
	fmt.Println(n)

}

func HashBucket(word string, buckets int32) int32 {
	letter := rune(word[0])
	bucket := letter % buckets
	return bucket
}
