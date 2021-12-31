package main

import (
	"fmt"
	"os"
	"sync"
	"words/processing"
)

func main() {
	in := make(chan string, 10)
	out := make(chan string, 10)

	go processing.DictionaryMatches(in, out)

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		for _, word := range os.Args[1:] {
			in <- word
		}
		close(in)
		wg.Done()
	}()

	go func() {
		for word := range out {
			fmt.Printf("%q\n", word)
		}
		wg.Done()
	}()

	wg.Wait()
}
