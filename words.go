package main

import (
	"fmt"
	"os"
	"words/processing"
)

func main() {
	ch1 := make(chan string, 10)
	ch2 := make(chan string, 10)
	ch3 := make(chan string, 10)
	ch4 := make(chan string, 10)

	go processing.WordSubsets(os.Args[1], ch1)
	go processing.DictionaryMatches(ch1, ch2)
	go processing.Deduplicate(ch2, ch3)
	go processing.Sort(ch3, ch4)

	for word := range ch4 {
		fmt.Printf("%s\n", word)
	}
}
