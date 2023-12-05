package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func Part1(fn string) {
	file, err := os.Open(fn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		runes := []rune(line)
		l, r := 0, len(runes)-1
		for ; l <= r && !unicode.IsDigit(runes[l]); l++ {
		}
		for ; l <= r && !unicode.IsDigit(runes[r]); r-- {
		}
		sum += (int(runes[l])-int('0'))*10 + (int(runes[r]) - int('0'))
	}
	fmt.Println(sum)
}

func ParseDigit(runes []rune, pos int) int {
	if unicode.IsDigit(runes[pos]) {
		return int(runes[pos]) - int('0')
	}
	l := len(runes)
	for i, word := range []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"} {
		chars := len(word)
		if pos+chars <= l && string(runes[pos:pos+chars]) == word {
			return i + 1
		}
	}
	return -1
}

func Part2(fn string) {
	file, err := os.Open(fn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		runes := []rune(line)
		l, r := 0, len(runes)-1
		for ; l <= r; l++ {
			if digit := ParseDigit(runes, l); digit > 0 {
				sum += digit * 10
				break
			}
		}
		for ; l <= r; r-- {
			if digit := ParseDigit(runes, r); digit > 0 {
				sum += digit
				break
			}
		}
	}
	fmt.Println(sum)
}

func main() {
	part, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	filename := os.Args[2]
	switch part {
	case 1:
		Part1(filename)
	case 2:
		Part2(filename)
	default:
		fmt.Fprintf(os.Stderr, "Error: %v\n", "Invalid part")
		os.Exit(1)
	}
}
