package main

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
)

type Day01 struct {
	sampleSize int
}

func (d Day01) Solve(scanner *bufio.Scanner) {
	var err error
	var depths = make([]int, d.sampleSize)
	var sum_a, sum_b, numIncreases, numDepths int

	for scanner.Scan() {
		i := numDepths % d.sampleSize
		depths[i], err = strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		sum_a = sum_b
		sum_b = 0
		for _, depth := range depths {
			sum_b += depth
		}

		if numDepths < d.sampleSize {
			numDepths++
			continue
		}
		numDepths++

		if sum_b > sum_a {
			numIncreases++
		}
	}

	fmt.Printf("The number of depth increases for a sample size of %d is: %d\n", d.sampleSize, numIncreases)
}
