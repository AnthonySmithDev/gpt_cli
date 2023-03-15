package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/AnthonySmithDev/gpt/prompt"
)

// Example:
// how to use bufio.Scanner to read multiple lines of text in golang

func main() {
	var (
		ask  bool
		code bool
	)

	flag.BoolVar(&ask, "a", false, "Ask from GPT")
	flag.BoolVar(&ask, "ask", false, "Ask from GPT")

	flag.BoolVar(&code, "c", false, "Completion from GPT")
	flag.BoolVar(&code, "code", false, "Completion from GPT")

	flag.Parse()

	lines, err := getInput()
	if err != nil {
		panic(err)
	}

	text := strings.Join(lines, "\n")
	if text == "" {
		panic(fmt.Errorf("No input text"))
	}
	text = text + ". "

	switch {
	case ask:
		err = prompt.AskCompletion(text)
		if err != nil {
			panic(err)
		}
	case code:
		err = prompt.CodeCompletion(text)
		if err != nil {
			panic(err)
		}
	default:
		fmt.Println(text)
	}

}

func getInput() ([]string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for {
		scanner.Scan()
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		lines = append(lines, line)
	}
	err := scanner.Err()
	if err != nil {
		return nil, err
	}
	return lines, nil
}
