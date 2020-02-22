package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// binarySearch returns the index of item from sorted list or -1 if not found.
func binarySearch(item int, list []int) int {
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := int((low + high) / 2)
		guess := list[mid]
		if guess == item {
			return mid
		}
		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

func stringToInts(ss []string) []int {
	var is []int
	for _, s := range ss {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		is = append(is, i)
	}
	return is
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Fprintf(os.Stderr, "usage: %s <int-to-find> <sorted-list-of-ints>\n", os.Args[0])
		os.Exit(1)
	}

	item, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	list := stringToInts(os.Args[2:])
	fmt.Println(binarySearch(item, list))
}
