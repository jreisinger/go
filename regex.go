package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <regex> <file>\n", os.Args[0])
		os.Exit(1)
	}

	re, err := regexp.Compile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	f, err := os.Open(os.Args[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	//scanner.Split(bufio.ScanLines) // default
	for scanner.Scan() {
		line := scanner.Text()
		if re.MatchString(line) {
			fmt.Println(line)
		}
	}
}
