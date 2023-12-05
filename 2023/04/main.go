package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
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

	spaces := regexp.MustCompile(`\s+`)

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		line = spaces.ReplaceAllString(line, " ")

		parts := strings.Split(line, ": ")
		numberParts := strings.Split(parts[1], " | ")
		winningNumbers := strings.Split(numberParts[0], " ")
		cardNumbers := strings.Split(numberParts[1], " ")

		matches := 0
		for _, number := range cardNumbers {
			for _, winningNumber := range winningNumbers {
				if number == winningNumber {
					matches++
				}
			}
		}
		if matches > 0 {
			points := 1 << (matches - 1)
			total += points
		}
	}
	fmt.Println(total)
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

	spaces := regexp.MustCompile(`\s+`)

	cards := 0
	copies := map[int]int{}
	for scanner.Scan() {
		line := scanner.Text()
		line = spaces.ReplaceAllString(line, " ")

		parts := strings.Split(line, ": ")
		cardParts := strings.Split(parts[0], " ")
		cardNumber, _ := strconv.Atoi(cardParts[1])
		numberParts := strings.Split(parts[1], " | ")
		winningNumbers := strings.Split(numberParts[0], " ")
		cardNumbers := strings.Split(numberParts[1], " ")

		copies[cardNumber]++
		cards += copies[cardNumber]
		matches := 0
		for _, number := range cardNumbers {
			for _, winningNumber := range winningNumbers {
				if number == winningNumber {
					matches++
				}
			}
		}
		for i := cardNumber + 1; i <= cardNumber+matches; i++ {
			copies[i] += copies[cardNumber]
		}
	}
	fmt.Println(cards)
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
