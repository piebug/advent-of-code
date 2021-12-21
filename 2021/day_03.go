package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Day03 struct{}

func (d Day03) Solve(scanner *bufio.Scanner) {
	var zeroes, ones []string
	var bitCounter []int
	var total int

	for scanner.Scan() {
		bitString := scanner.Text()
		bits := strings.Split(bitString, "")
		total++

		if bitString[0] == '1' {
			ones = append(ones, bitString)
		} else {
			zeroes = append(zeroes, bitString)
		}

		for i, bit := range bits {
			if bitCounter == nil || i >= len(bitCounter) {
				bitCounter = append(bitCounter, 0)
			}
			if bit == "1" {
				bitCounter[i]++
			}
		}
	}

	var oxygen, co2 string

	if len(zeroes) > len(ones) {
		oxygen = findLifeSupportRating(zeroes, false)
		co2 = findLifeSupportRating(ones, true)
	} else {
		oxygen = findLifeSupportRating(ones, false)
		co2 = findLifeSupportRating(zeroes, true)
	}

	fmt.Printf("The life support rating is: %d\n", parseResults(oxygen, co2))

	var gamma, epsilon string
	half := total / 2

	for _, count := range bitCounter {
		if count > half {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	fmt.Printf("The power consumption is: %d\n", parseResults(gamma, epsilon))
}

func findLifeSupportRating(bits []string, inverse bool) string {
	rating := string(bits[0][0])
	columns := len(bits[0])
	// Starts at one because we filter out the first column during the initial scan
	for i := 1; i < columns; i++ {
		bigger, smaller := filterBits(bits, i)
		if inverse && len(smaller) != 0 {
			rating += string(smaller[0][i])
			bits = smaller
		} else {
			rating += string(bigger[0][i])
			bits = bigger
		}
		if len(bits) == 1 {
			return bits[0]
		}
	}
	return rating
}

func filterBits(bits []string, column int) (biggerBits []string, smallerBits []string) {
	var zeroBits, oneBits []string
	for _, bit := range bits {
		if bit[column] == '1' {
			oneBits = append(oneBits, bit)
		} else {
			zeroBits = append(zeroBits, bit)
		}
	}

	if len(zeroBits) > len(oneBits) {
		biggerBits = zeroBits
		smallerBits = oneBits
	} else {
		biggerBits = oneBits
		smallerBits = zeroBits
	}
	return
}

func parseResults(var1 string, var2 string) int64 {
	var1Int, _ := strconv.ParseInt(var1, 2, 64)
	var2Int, _ := strconv.ParseInt(var2, 2, 64)

	return var1Int * var2Int
}
