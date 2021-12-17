package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func day01() {
	file, err := os.Open("data/day_01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Scanner by default splits at newline
	scanner := bufio.NewScanner(file)

	var depth_a, depth_b, numDeeper int

	for scanner.Scan() {
		depth_a = depth_b
		depth_b, err = strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		if depth_a == 0 {
			continue
		}

		if depth_b > depth_a {
			numDeeper++
		}
	}

	fmt.Printf("num deeper: %d", numDeeper)
}
