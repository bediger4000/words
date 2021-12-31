package processing

import "sort"

func Sort(in, out chan string) {
	var words []string
	for word := range in {
		words = append(words, word)
	}
	sort.Strings(words)
	for i := range words {
		out <- words[i]
	}
	close(out)
}
