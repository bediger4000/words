package processing

import "fmt"

func Tee(in, out1, out2 chan string) {
	for word := range in {
		out1 <- word
		out2 <- word
	}

	close(out1)
	close(out2)
}

func TeeOut(in chan string) {
	for word := range in {
		fmt.Printf("%s\n", word)
	}
}
