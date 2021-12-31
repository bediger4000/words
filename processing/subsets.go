package processing

func WordSubsets(word string, out chan string) {
	runes := []rune(word)
	realfindwords(runes, out)
	close(out)
}

func realfindwords(runes []rune, out chan string) {
	if len(runes) < 3 {
		return
	}

	out <- string(runes)

	next := make([]rune, len(runes)-1)

	for i := range runes {
		j := 0
		for k := range runes {
			if k == i {
				continue
			}
			next[j] = runes[k]
			j++
		}
		realfindwords(next, out)
	}
}
