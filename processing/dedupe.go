package processing

func Deduplicate(in, out chan string) {
	seen := make(map[string]bool)

	for word := range in {
		if seen[word] {
			continue
		}
		seen[word] = true
		out <- word
	}

	close(out)
}
