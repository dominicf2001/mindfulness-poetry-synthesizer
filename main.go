package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
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

		i := poem.currentPos
		for poem.body[i] != ')' {
			i++
		}

		poem.currentPrompt = poem.body[poem.currentPos : i+1]
	}
	poem.promptsRemaining--
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

	poem := poems[rand.Intn(len(poems))]

	fmt.Println("Use only one word in your responses.\n")
	for poem.promptsRemaining > 0 {
		fmt.Printf("%s: ", poem.currentPrompt)

		var input string
		// prompt for user input
		fmt.Scanln(&input)

		poem.insertInput(input)
		fmt.Println()
	}

	fmt.Println(poem.body)

	f.Close()
}
