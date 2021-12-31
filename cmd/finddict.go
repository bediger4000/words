package main

import (
	"fmt"
	"os"
	"words/processing"
)

func main() {

	dict := processing.MatchingDictionary()

	fmt.Printf("%d keys in dictionary\n", len(dict))

	for _, word := range os.Args[1:] {
		if matches, ok := dict[processing.WordKey(word)]; ok {
			fmt.Printf("%s: %v\n", word, matches)
		}
	}
}
