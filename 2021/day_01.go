package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func day01(sampleSize int) {
	file, err := os.Open("data/day_01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Scanner by default splits at newline
	scanner := bufio.NewScanner(file)

	var depths = make([]int, sampleSize)
	var sum_a, sum_b, numIncreases, numDepths int

	for scanner.Scan() {
		i := numDepths % sampleSize
		depths[i], err = strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		sum_a = sum_b
		sum_b = 0
		for _, depth := range depths {
			sum_b += depth
		}

		if numDepths < sampleSize {
			numDepths++
			continue
		}
		numDepths++

		if sum_b > sum_a {
			numIncreases++
		}
	}

	fmt.Printf("The number of depth increases for a sample size of %d is: %d\n", sampleSize, numIncreases)
}
