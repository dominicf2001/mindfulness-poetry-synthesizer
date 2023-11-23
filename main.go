package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type poem struct {
	title         string
	body          string
	currentPrompt string
	currentPos    int
}

func (poem *poem) displayPrompt() {
	fmt.Printf("Enter %s\n", poem.currentPrompt)
}

func (poem *poem) insertInput() {

}

func (poem *poem) next() {

}

func newPoem(title string, body string) *poem {

	return &poem{title: title, body: body}
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

	f.Close()
}
