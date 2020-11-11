package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

func grepStdIn(input string, v bool, i bool, g bool) {
	stdInReader := bufio.NewReader(os.Stdin)

	grep(stdInReader, input, v, i, g)
}

func grepFile(fileAddress, input string, v bool, i bool, g bool) {
	fileReader, err := os.Open(fileAddress)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer fileReader.Close()

	grep(fileReader, input, v, i, g)
}

func grep(reader io.Reader, input string, v bool, i bool, g bool) {
	scanner := bufio.NewScanner(reader)
	rgx := regexp.MustCompile(input)

	for scanner.Scan() {
		line := scanner.Text()

		if i {
			input = strings.ToLower(input)
			line = strings.ToLower(line)
		}

		if g {
			match := rgx.MatchString(line)

			if match && !v {
				fmt.Println(line)
			} else if !match && v {
				fmt.Println(line)
			}
		} else {
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
}

func main() {
	fPtr := flag.String("f", "Stdin", "Take patterns from file")
	vPtr := flag.Bool("v", false, "Select non-matching lines")
	iPtr := flag.Bool("i", false, "Ignore case distinctions in patterns and data")
	gPtr := flag.Bool("g", false, "PATTERNS are basic regular expressions")

	flag.Parse()

	input := flag.Arg(0)

	if fPtr != nil && *fPtr != "Stdin" {
		grepFile(*fPtr, input, *vPtr, *iPtr, *gPtr)
	} else {
		grepStdIn(input, *vPtr, *iPtr, *gPtr)
	}
}
