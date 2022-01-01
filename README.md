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
[Knight's Tour](https://github.com/bediger4000/knights-tour)
[N Queens](https://github.com/bediger4000/nqueens)
and
[Sudoku](https://github.com/bediger4000/sudoku-solver-2)
by having a goroutine run a recursive function.
When recursion terminates with a solution,
the goroutine writes the solution to a channel.
The main goroutine reads solutions and possibly
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

This package makes a very easy main function:

```go
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
```

`package processing` isn't at all complete.
It only hints at a useful set of functions.

A useful package would name an interface for the functions:

```go
type Filter interface {
	Run(in, out chan string, arguments ...interface{})
}
```

At this point, I can't tell if naming "source" and "sink" interfaces
would help.
I wrote such functions, `processing.WordSubsets` and `processing.TeeOut`.

The obvious functions missing from the package are named file inputs and outputs,
and regular expression matching.
Other Unix shell commands suggest further functionality: `tr`, `uniq`,
and counters like `wc`.

Once you've got a well-designed text-processing package,
the next step is an interpreter that creates channels as necessary,
and strings together goroutines running functions.
This sounds rather like a text-processing shell,
along the lines of [jq](https://stedolan.github.io/jq/),
which handles JSON, not plain text,
or [GNU datamash](https://www.gnu.org/software/datamash/).

Tenatively, the command language would look a lot like traditional
Unix shell "pipelines".
It looks like if you don't remember the past,
you are doomed to repeat it.
