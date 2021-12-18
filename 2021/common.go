package main

import (
	"bufio"
	"log"
	"os"
)

type AdventDay interface {
	Solve(scanner *bufio.Scanner)
}

func ScanSolve(filename string, day AdventDay) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Scanner by default splits at newline
	scanner := bufio.NewScanner(file)
	day.Solve(scanner)
}
