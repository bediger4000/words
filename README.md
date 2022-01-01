# Shell-style text processing in Go

I saw a word game where the puzzle gives you six letters.
By means of a clever user interface, you construct words from 3 or more
of the letters.
You use the words to fill in a cross-word puzzle style,
cross-linked arrangement of words.

[Wordscapes](https://play.google.com/store/apps/details?id=com.peoplefun.wordcross&hl=en&gl=us)

I wrote a program to help find those words.

## Design

I've written backtracking solutions to
[Knight's Tour]()
[N Queens]()
and
[Sudoku]()
by having a goroutine running a recursive function.
When recursion terminates with a solution,
the goroutine writes the solution to a channel.
The main goroutine reads solutions and possible
eliminates duplicates.

I followed that pattern for this problem.
I wrote a program that generated subsets of the letters of the solution words,
and wrote them to a channel.
Then I wrote a program that alphabetizes those letter subsets,
and finds dictionary (`/usr/share/dict/words`) words that get spelled with a subset.
It occurred to me to change those programs to functions that accept a
Golang "in" channel and an "out" channel, run the functions in goroutines,
and have the main program deduplicate solution words before output.
I experienced a bug, and ended up writing a "tee" function that wrote
words from its input channel to stdout before writing those words to its output channel.
At that point, I had a pattern to follow, so I re-wrote all the functions
to have "in" and "out" channels.
`func main()` creates channels of string to link functions,
then runs the functions each in their own goroutine.

I ended up with `package processing`

## package processing

```go
package processing // import "words/processing"

func CreateDictionary() map[string]bool
func MatchingDictionary() map[string][]string

func Deduplicate(in, out chan string)
func DictionaryCheck(in, out chan string)
func DictionaryMatches(in, out chan string)
func Sort(in, out chan string)
func WordKey(word string) string
func WordSubsets(word string, out chan string)

func Tee(in, out1, out2 chan string)
func TeeOut(in chan string)
```
