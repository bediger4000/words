package main

import (
	"os"
	"words/processing"
)

func main() {
	ch1 := make(chan string, 10)
	ch2 := make(chan string, 10)
	ch3 := make(chan string, 10)
	ch4 := make(chan string, 10)

	go processing.WordSubsets(os.Args[1], ch1) // generate subsets of characters
	go processing.DictionaryMatches(ch1, ch2)  // dictionary words from subsets
	go processing.Deduplicate(ch2, ch3)        // remove duplicate dictionary words
	go processing.Sort(ch3, ch4)               // sort dictionary words alphabetically

	processing.TeeOut(ch4)
}
