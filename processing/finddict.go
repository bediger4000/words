package processing

import (
	"bufio"
	"log"
	"os"
	"sort"
)

func DictionaryMatches(in, out chan string) {
	possibilities := MatchingDictionary()
	for word := range in {
		if matches, ok := possibilities[WordKey(word)]; ok {
			for _, match := range matches {
				out <- match
			}
		}
	}
	close(out)
}

func MatchingDictionary() map[string][]string {
	fin, err := os.Open("/usr/share/dict/words")
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()

	dict := make(map[string][]string)

	scanner := bufio.NewScanner(fin)
	for scanner.Scan() {
		line := scanner.Text()
		key := WordKey(line)
		dict[key] = append(dict[key], line)
	}

	return dict
}

func WordKey(word string) string {
	runes := []rune(word)
	sort.Sort(runeSlice(runes))
	return string(runes)
}

type runeSlice []rune

func (r runeSlice) Len() int {
	return len(r)
}

func (r runeSlice) Less(i, j int) bool {
	return r[i] < r[j]
}

func (r runeSlice) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
