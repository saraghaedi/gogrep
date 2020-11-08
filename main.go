package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func grepStdIn(input string, v bool, i bool) {
	stdInReader := bufio.NewReader(os.Stdin)

	grep(stdInReader, input, v, i)
}

func grepFile(fileAddress, input string, v bool, i bool) {
	fileReader, err := os.Open(fileAddress)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer fileReader.Close()

	grep(fileReader, input, v, i)
}

func grep(reader io.Reader, input string, v bool, i bool) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()

		if i {
			input = strings.ToLower(input)
			line = strings.ToLower(line)
		}

		if !v {
			if strings.Contains(line, input) {
				fmt.Println(line)
			}
		} else {
			if !strings.Contains(line, input) {
				fmt.Println(line)
			}
		}
	}
}

func main() {
	fPtr := flag.String("f", "Stdin", "Take patterns from file")
	vPtr := flag.Bool("v", false, "Select non-matching lines")
	iPtr := flag.Bool("i", false, "Ignore case distinctions in patterns and data")

	flag.Parse()

	input := flag.Arg(0)

	if fPtr != nil && *fPtr != "Stdin" {
		grepFile(*fPtr, input, *vPtr, *iPtr)
	} else {
		grepStdIn(input, *vPtr, *iPtr)
	}
}
