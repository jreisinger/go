// https://yourbasic.org/golang/write-command-line-application/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	log.SetPrefix("grep: ")
	log.SetFlags(0) // no extra info in log messages

	if len(os.Args) != 3 {
		log.Fatalf("Usage: %v PATTERN FILE\n", os.Args[0])
	}

	pattern, err := regexp.Compile(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	file, err := os.Open(os.Args[2])
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	//scanner.Split(bufio.ScanLines) // default
	for scanner.Scan() {
		line := scanner.Text()
		if pattern.MatchString(line) {
			fmt.Println(line)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
}
