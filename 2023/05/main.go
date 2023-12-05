package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Translate1(seeds []int, seedMaps [][3]int) (result []int) {
	result = make([]int, len(seeds))
SeedLoop:
	for i, seed := range seeds {
		for _, mapping := range seedMaps {
			if seed >= mapping[1] && seed < mapping[1]+mapping[2] {
				result[i] = mapping[0] + seed - mapping[1]
				continue SeedLoop
			}
		}
		result[i] = seed
	}
	return
}

func Translate2(seeds []int, seedMaps [][3]int) (result []int) {
	result = make([]int, 0)
	heap := make([]int, 0)
	heap = append(heap, seeds...)
SeedLoop:
	for i := 0; i < len(heap); i += 2 {
		seedFrom := heap[i]
		seedTo := seedFrom + heap[i+1] - 1
		for _, mapping := range seedMaps {
			mappingFrom := mapping[1]
			mappingTo := mappingFrom + mapping[2] - 1

			if seedFrom >= mappingFrom && seedTo <= mappingTo {
				result = append(result, mapping[0]+seedFrom-mappingFrom, seedTo-seedFrom+1)
				continue SeedLoop
			}
			if seedFrom < mappingFrom && seedTo > mappingTo {
				result = append(result, mapping[0], mappingTo-mappingFrom+1)
				heap = append(heap, seedFrom, mappingFrom-seedFrom)
				heap = append(heap, mappingTo+1, seedTo-mappingTo)
				continue SeedLoop
			}
			if seedFrom >= mappingFrom && seedFrom <= mappingTo && seedTo > mappingTo {
				result = append(result, mapping[0]+seedFrom-mappingFrom, mappingTo-seedFrom+1)
				heap = append(heap, mappingTo+1, seedTo-mappingTo)
				continue SeedLoop
			}
			if seedFrom < mappingFrom && seedTo >= mappingFrom && seedTo <= mappingTo {
				result = append(result, mapping[0], seedTo-mappingFrom+1)
				heap = append(heap, seedFrom, mappingFrom-seedFrom)
				continue SeedLoop
			}
		}
		result = append(result, seedFrom, seedTo-seedFrom+1)
	}
	return
}

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

	seeds := make([]int, 0)
	seedMaps := make([][3]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		line = spaces.ReplaceAllString(line, " ")
		if len(seeds) == 0 {
			seedParts := strings.Split(line, ": ")
			numberParts := strings.Split(seedParts[1], " ")
			for _, number := range numberParts {
				seedNumber, _ := strconv.Atoi(number)
				seeds = append(seeds, seedNumber)
			}
		} else {
			if len(line) == 0 {
				continue
			}
			if strings.HasSuffix(line, " map:") {
				if len(seedMaps) > 0 {
					seeds = Translate1(seeds, seedMaps)
					seedMaps = make([][3]int, 0)
				}
			} else {
				mapParts := strings.Split(line, " ")
				mapping := [3]int{}
				for i, number := range mapParts {
					mapping[i], _ = strconv.Atoi(number)
				}
				seedMaps = append(seedMaps, mapping)
			}
		}
	}
	seeds = Translate1(seeds, seedMaps)
	min := seeds[0]
	for _, seed := range seeds {
		if seed < min {
			min = seed
		}
	}
	fmt.Println(min)
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
	seeds := make([]int, 0)
	seedMaps := make([][3]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		line = spaces.ReplaceAllString(line, " ")
		if len(seeds) == 0 {
			seedParts := strings.Split(line, ": ")
			numberParts := strings.Split(seedParts[1], " ")
			for _, number := range numberParts {
				seedNumber, _ := strconv.Atoi(number)
				seeds = append(seeds, seedNumber)
			}
		} else {
			if len(line) == 0 {
				continue
			}
			if strings.HasSuffix(line, " map:") {
				if len(seedMaps) > 0 {
					seeds = Translate2(seeds, seedMaps)
					seedMaps = make([][3]int, 0)
				}
			} else {
				mapParts := strings.Split(line, " ")
				mapping := [3]int{}
				for i, number := range mapParts {
					mapping[i], _ = strconv.Atoi(number)
				}
				seedMaps = append(seedMaps, mapping)
			}
		}
	}
	seeds = Translate2(seeds, seedMaps)
	min := seeds[0]
	for i := 2; i < len(seeds); i += 2 {
		if seeds[i] < min {
			min = seeds[i]
		}
	}
	fmt.Println(min)
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
