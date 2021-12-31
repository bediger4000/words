package main

import (
	"fmt"
	"os"

	"words/processing"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go processing.WordSubsets(os.Args[1], ch1)
	go processing.Deduplicate(ch1, ch2)

	for word := range ch2 {
		fmt.Printf("%q\n", word)
	}
}
