package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	s "strings"
)

type poem struct {
	title            string
	body             string
	currentPrompt    string
	currentPos       int
	promptsRemaining int
}

func (poem *poem) displayPrompt() {
	fmt.Printf("Enter %s\n", poem.currentPrompt)
}

func (poem *poem) insertInput(input string) {
	poem.body = s.Replace(poem.body, poem.currentPrompt, input, len(input))

	poem.next()
}

func (poem *poem) next() {
	if poem.promptsRemaining > 1 {
		for poem.body[poem.currentPos] != '(' {
			poem.currentPos++
		}
		poem.currentPos++
	}

	i := poem.currentPos
	for poem.body[i] != ')' {
		i++
	}

	poem.promptsRemaining--
	poem.currentPrompt = poem.body[poem.currentPos:i]
}

func newPoem(title string, body string) *poem {
	p := &poem{title: title, body: body}

	// p.body is not inserted into,
	// thus make up for the decrement in next() by adding 1
	p.promptsRemaining = s.Count(p.body, "(") + 1
	p.next()

	return p
}

func main() {
	f, err := os.Open("./poems")

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)

	poems := make([]*poem, 0)

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		if line[0] == '#' {
			title := line
			body, err := reader.ReadString('-')

			if err != nil {
				panic(err)
			}

			poems = append(poems, newPoem(title, body))
		}
	}

	for poems[0].promptsRemaining > 0 {
		fmt.Println(poems[0].currentPrompt)
		poems[0].next()
	}

	f.Close()
}
