package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func printFileLong(fileName string) {
	// open a file for reading
	file, err := os.Open(fileName) // file is *File
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// get the file size
	fileInfo, err := file.Stat() // func (*File) Stat -> one of *File's methods
	if err != nil {
		log.Fatal(err)
	}

	// read the file
	bs := make([]byte, fileInfo.Size())
	_, err = file.Read(bs)
	if err != nil {
		log.Fatal(err)
	}

	// print the read data
	str := string(bs)
	fmt.Println(str)
}

func printFileShort(fileName string) {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	str := string(bs)
	fmt.Println(str)
}

func main() {
	fileName := os.Args[1]

	printFileLong(fileName)
	printFileShort(fileName)
}
