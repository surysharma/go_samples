package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func closeFile(f *os.File) {
	err := f.Close()
	if err != nil {
		log.Fatal("Unable to close the file", err)
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: filecmp <file1> <file2>")
		os.Exit(1)
	}

	file1 := os.Args[1]
	file2 := os.Args[2]

	f1, err := os.Open(file1)
	if err != nil {
		fmt.Println("Error opening", file1, ":", err)
		os.Exit(1)
	}
	defer closeFile(f1)

	f2, err := os.Open(file2)
	if err != nil {
		fmt.Println("Error opening", file2, ":", err)
		os.Exit(1)
	}
	defer closeFile(f2)

	scanner1 := bufio.NewScanner(f1)
	scanner2 := bufio.NewScanner(f2)
	line := 1
	for scanner1.Scan() && scanner2.Scan() {
		if scanner1.Text() != scanner2.Text() {
			fmt.Printf("Line %d:\n< %s\n> %s\n", line, scanner1.Text(), scanner2.Text())
			line++
		}
	}

	if scanner1.Err() != nil {
		fmt.Println("Error reading", file1, ":", scanner1.Err())
		os.Exit(1)
	}

	if scanner2.Err() != nil {
		fmt.Println("Error reading", file2, ":", scanner2.Err())
		os.Exit(1)
	}

	if line > 1 {
		fmt.Println("Files are different")
	} else {
		fmt.Println("Files are the same")
	}
}
