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

	elfMap := make(map[int]int)

	for scanner.Scan() {

		if scanner.Text() == "" {
			elfMap[currentElf] = currentElfCals
			currentElf += 1
			currentElfCals = 0
		} else {
			i, _ := strconv.Atoi(scanner.Text())
			currentElfCals += i
		}
	}

	type keyValue struct {
		Key   int
		Value int
	}

	var sortedSet []keyValue
	for k, v := range elfMap {
		sortedSet = append(sortedSet, keyValue{k, v})
	}

	sort.Slice(sortedSet, func(i, j int) bool {
		return sortedSet[i].Value > sortedSet[j].Value
	})

	for _, kv := range sortedSet {
		fmt.Printf("%d, %d\n", kv.Key, kv.Value)
	}

	topThree := sortedSet[0:3]
	fmt.Printf("Top 3: %v \n", topThree)

	fmt.Println(sortedSet[0].Value + sortedSet[1].Value + sortedSet[2].Value)

}
