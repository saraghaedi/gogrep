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

func grepStdIn(input string) {
	stdInReader := bufio.NewReader(os.Stdin)

	grep(stdInReader, input)
}

func grepFile(fileAddress, input string) {
	fileReader, err := os.Open(fileAddress)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer fileReader.Close()

	grep(fileReader, input)
}

func grep(reader io.Reader, input string) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, input) {
			fmt.Println(line)
		}
	}
}

func main() {
	fPtr := flag.String("f", "Stdin", "Take patterns from file")

	flag.Parse()

	input := flag.Arg(0)

	if fPtr != nil && *fPtr != "Stdin" {
		grepFile(*fPtr, input)
	} else {
		grepStdIn(input)
	}
}
