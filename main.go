package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

	filename := "input"
	inputFile, _ := os.Open(filename)
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	currentElfCals := 0
	currentElf := 0

	// Slice of my array looks like
	type elfCals struct {
		Elf  int
		Cals int
	}

	// Array
	var elfMap []elfCals

	// Each line of file
	for scanner.Scan() {

		// if it's a blank line, bank the cals, increment the elf
		if scanner.Text() == "" {
			elfMap = append(elfMap, elfCals{currentElf, currentElfCals})
			currentElf += 1
			currentElfCals = 0
		} else {
			// if it's not a new elf, add the cals to the running total
			i, _ := strconv.Atoi(scanner.Text())
			currentElfCals += i
		}
	}

	// sort the slices
	sort.Slice(elfMap, func(i, j int) bool {
		return elfMap[i].Cals > elfMap[j].Cals
	})

	// print the array for fun
	for _, kv := range elfMap {
		fmt.Printf("%d, %d\n", kv.Elf, kv.Cals)
	}

	// list the top three
	topThree := elfMap[0:3]
	fmt.Printf("Top 3: %v \n", topThree)

	// add their cals together
	fmt.Println(elfMap[0].Cals + elfMap[1].Cals + elfMap[2].Cals)

}
