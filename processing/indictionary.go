package processing

import (
	"bufio"
	"log"
	"os"
)

func DictionaryCheck(in, out chan string) {
	found := CreateDictionary()
	for word := range in {
		if found[word] {
			out <- word
		}
	}
	close(out)
}

// CreateDictionary reads in /usr/share/dict/words
// returning a map with words in dictionary as keys,
// true as values.
func CreateDictionary() map[string]bool {
	fin, err := os.Open("/usr/share/dict/words")
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()

	dict := make(map[string]bool)

	scanner := bufio.NewScanner(fin)
	for scanner.Scan() {
		line := scanner.Text()
		dict[line] = true
	}

	// fmt.Fprintf(os.Stderr, "Found %d words in dictionary\n", len(dict))

	return dict
}
